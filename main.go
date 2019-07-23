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
	"flag"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/r3labs/sse"
)

const (
	ctxHubKey = "__hub_key__"

	KindSubmission = "rs"
)

var (
	staticFileServerAddress = flag.String("static", "http://localhost:1234", "The path to the static file server")
	listenAddress           = flag.String("listen", ":8222", "The address to listen too")
)

func sseFunnelBytesSub(c chan<- []byte) {
	for {
		client := sse.NewClient("http://stream.pushshift.io?type=submissions")
		client.OnDisconnect(func(c *sse.Client) {
			log.Printf("We got discconected, but we are trying to reconnect\n")
		})

		err := client.Subscribe("rs", func(msg *sse.Event) { c <- msg.Data })
		if err != nil {
			log.Printf("We got an error from subscribe: %s\n", err.Error())
		}
	}
}

func main() {
	flag.Parse()

	hub := newHub()
	go hub.run()
	go sseFunnelBytesSub(hub.broadcast)

	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.WithValue(ctxHubKey, hub))

	r.Get("/ws", serveWs)

	pxu, err := url.Parse(*staticFileServerAddress)
	if err != nil {
		log.Fatalf("Invalid static fileserver error: %s\n", err.Error())
	}
	px := httputil.NewSingleHostReverseProxy(pxu)
	r.NotFound(px.ServeHTTP)

	if err := http.ListenAndServe(*listenAddress, r); err != nil {
		log.Fatal(err)
	}
}
