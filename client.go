// Copyright (c) 2019 Teodor Spæren
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package main

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait = 10 * time.Second

	pongWait = 60 * time.Second

	pingPeriod = (pongWait * 9) / 10

	maxMessageSize = 4096
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
}

type SubFilter func(*RedditSubmissionRefined) bool

type Client struct {
	hub *Hub

	conn *websocket.Conn

	send chan []byte

	// A filter. Any one of these return false, we don't want the message
	filters []SubFilter
}

func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()

	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		mt, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v\n", err)
			}
			break
		}
		log.Printf("Got message of type: %d and %s\n", mt, message)
	}
}

func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write([]byte{'\n'})
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				log.Printf("Error in write: %s\n", err.Error())
				return
			}

		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func SubredditNameFilter(subs []string) SubFilter {
	wewant := make(map[string]struct{})
	for _, s := range subs {
		// if there was no query string specified, we will get a single array element with "". We want
		// to filter that out, as it makes the checking for no filters much easier
		wewant[strings.ToLower(s)] = struct{}{}
	}

	return func(r *RedditSubmissionRefined) bool {
		_, ok := wewant[strings.ToLower(r.Subreddit)]
		return ok
	}
}

func AgeRatingFilter(nsfw bool) SubFilter {
	return func(r *RedditSubmissionRefined) bool {
		return r.Over18 == nsfw
	}
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	hub := r.Context().Value(ctxHubKey).(*Hub)

	qs := r.URL.Query()

	fs := make([]SubFilter, 0)

	// subreddit filter
	rSubs := qs.Get("subreddits")
	if rSubs != "" {
		fs = append(fs, SubredditNameFilter(strings.Split(rSubs, ",")))
	}

	// nsfw filter
	rRating := strings.ToLower(qs.Get("rating"))
	if rRating != "" {
		var nsfw bool

		if rRating == "nsfw" {
			nsfw = true
		} else if rRating == "sfw" {
			nsfw = false
		} else {
			log.Printf("We got a weird rating value!: %s\n", rRating)
			http.Error(w, "bad rating value", http.StatusBadRequest)
			return
		}
		fs = append(fs, AgeRatingFilter(nsfw))
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256), filters: fs}
	client.hub.register <- client

	go client.writePump()
	go client.readPump()
}
