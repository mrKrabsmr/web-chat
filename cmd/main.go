package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients []websocket.Conn

type Message struct {
	Sender string `json:"sender"`
	Text   string `json:"text"`
}

func main() {
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Fatal(err)
		}

		clients = append(clients, *conn)

		for {
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}

			message := Message{
				Sender: conn.RemoteAddr().String(),
				Text:   string(msg),
			}

			jsonData, err := json.Marshal(message)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

			for _, client := range clients {
				if err = client.WriteMessage(msgType, jsonData); err != nil {
					return
				}

			}

		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", "text/html")
		if err := tmpl.Execute(w, nil); err != nil {
			http.Error(w, err.Error(), 500)
		}

	})

	http.ListenAndServe(":8080", nil)
}
