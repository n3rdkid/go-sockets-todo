package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var socketUpgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "We did it..... 🎉 ")
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

}
