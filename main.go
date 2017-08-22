package main

import (
	"net/http"

	"log"

	"easycast/currency"
	"easycast/server"
	"encoding/json"
	"strconv"
	"time"
)

var broadCaster *server.EasyCast

func init() {
	log.Println("Set config websocket")
	currency.InitCurrency()
	broadCaster = server.NewEasyCast(currency.UpdateCurrency, 1*time.Second, 5)
}

func simpleParam(r *http.Request, key string) (string, bool) {
	value, ok := r.Form[key]
	if !ok || len(value) != 1 {
		return "", false
	}

	return value[0], true
}

func connectWebSocketHandler(w http.ResponseWriter, r *http.Request) {
	ok := broadCaster.Subscribe(w, r)
	if !ok {
		log.Fatal("cannot subscribe")
	}
}

type Currencies struct {
	History *[]int `json:"history"`
}

type Error struct {
	Err string `json:"error"`
}

func getLastCurrency(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	w.Header().Set("Content-Type", "application/json")

	sizeStr, ok := simpleParam(r, "size")
	if !ok {
		json.NewEncoder(w).Encode(Error{"bad params"})
		return
	}

	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		json.NewEncoder(w).Encode(Error{"bad params"})
		return
	}

	if size != 10 {
		json.NewEncoder(w).Encode(Error{"bad size"})
		return
	}

	historyArray := currency.GetHistory(size)
	json.NewEncoder(w).Encode(Currencies{historyArray})
}

func main() {
	http.HandleFunc("/ws", connectWebSocketHandler)
	http.HandleFunc("/get", getLastCurrency)

	/* TODO nginx */
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	log.Println("Try listen 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
