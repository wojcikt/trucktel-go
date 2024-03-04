package main

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var upgrader websocket.Upgrader

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/ws", serveWs)

	server := http.Server{
		Handler:      router,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatalln(server.ListenAndServe())
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("upgrade failed: %v\n", err)
	}

	ws.WriteJSON(map[string]any{"x": 1})
}
