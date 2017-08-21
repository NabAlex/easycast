package main

import (
	"net/http"

	"log"

	"easycast/server"
	"time"
)

var broadCaster *server.EasyCast

func init() {
	log.Println("Set config websocket")
	broadCaster = server.NewEasyCast(1*time.Second, 5)
}

func connectWebSocketHandler(w http.ResponseWriter, r *http.Request) {
	ok := broadCaster.Subscribe(w, r)
	if !ok {
		log.Fatal("cannot subscribe")
	}
}

func main() {
	http.HandleFunc("/ws", connectWebSocketHandler)

	/* TODO nginx */
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	log.Println("Try listen 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
