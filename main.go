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
	ctxHubKey = "__hub_key__"

	KindSubmission = "rs"
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
