package main

import (
	"net/http"

	"github.com/gorilla/websocket"
	"log"
)

type Message struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

var upgrader websocket.Upgrader

func init() {
	log.Println("Set config websocket")
	upgrader = websocket.Upgrader{}
}

func connectWebSocketHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Try")

	wsocket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	defer wsocket.Close()

	err = wsocket.WriteJSON(struct {
		Str string `json:"hello"`
	}{ "hello" })

	if err != nil {
		log.Fatal(err)
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