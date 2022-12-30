package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/MarceloMPJR/gobingo/pkg/websocket"
)

var addr = flag.String("addr", ":8080", "http service address")

func main() {
	flag.Parse()

	hub := websocket.NewHub()
	go hub.Run()

	mux := http.NewServeMux()

	mux.Handle("/", http.HandlerFunc(serveHome))
	mux.Handle("/ws", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		websocket.ServeWs(hub, w, r)
	}))

	err := http.ListenAndServe(*addr, mux)
	if err != nil {
		panic(err)
	}
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println("request - home")
	http.ServeFile(w, r, "static/home.html")
}
