package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var socketUpgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	http.HandleFunc("/ws", handleWebSocket)
	fmt.Println("Server has started 🏃")
	http.ListenAndServe(":8080", nil)
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := socketUpgrade.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Failed to upgrade to socket", http.StatusBadRequest)
		return
	}
	defer conn.Close()

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read failed:", err)
			break
		}

		err = conn.WriteMessage(messageType, message)
		if err != nil {
			log.Println("write failed:", err)
			break
		}
	}

}
