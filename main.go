package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/r3labs/sse"
)

const (
	ctxRedditInKey = "__reddit_in_key__"
	ctxHubKey      = "__hub_key__"

	KindSubmission = "rs"
	KindComment    = "rc"
	KindKeepalive  = "keepalive"
)

func sseFunnelBytes(c chan<- []byte) {
	subs := make(chan *sse.Event)

	client := sse.NewClient("http://stream.pushshift.io?type=submissions")
	client.SubscribeChan("rs", subs)

	for {
		select {
		case k := <-subs:
			// c <- &RedditSSERaw{kind: KindSubmission, data: k.Data}
			c <- k.Data
		}
	}
}

func main() {

	hub := newHub()
	go hub.run()
	go sseFunnelBytes(hub.broadcast)

	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.WithValue(ctxHubKey, hub))

	r.Get("/ws", serveWs)

	// dev server
	pxu := &url.URL{
		Scheme: "http",
		Host:   "localhost:1234",
	}
	px := httputil.NewSingleHostReverseProxy(pxu)

	r.NotFound(px.ServeHTTP)

	if err := http.ListenAndServe(":8222", r); err != nil {
		log.Fatal(err)
	}
}
