package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		message := "Cannot upgrade http request to web socket"
		log.Fatal(message)
		http.Error(w, message, http.StatusInternalServerError)
		return
	}

	defer conn.Close()

	//Echos message back
	for {
		messageType, p, err := conn.ReadMessage()
		log.Println("Received Message")
		if err != nil {
			return
		}
		if err := conn.WriteMessage(messageType, p); err != nil {
			return
		}
	}
}

func main() {
	http.HandleFunc("/ws", handleWebSocket)
	http.ListenAndServe(":8080", nil)
}
