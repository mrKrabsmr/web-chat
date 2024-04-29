package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

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

func (c *Controller) Mailing() {
	for {
		message := Message{
			Sender: "ботяра",
			Text:   fmt.Sprintf("в чате онлайн участников: %d", len(clients)),
		}

		jsonData, err := json.Marshal(message)
			if err != nil {
				return
			}


		for _, client := range clients {
			client.WriteMessage(1, jsonData)
		}

		time.Sleep(time.Minute * 1)
	}
}

func (c *Controller) Chat(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		c.logger.Error(err)
		return
	}

	clients = append(clients, *conn)

	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			c.logger.Error(err)
			if _, ok := err.(*websocket.CloseError); ok {
				conn.Close()
				for i, client := range clients {
					if client.RemoteAddr().String() == conn.RemoteAddr().String() {
						clients = append(clients[:i], clients[i+1:]...)
					}
				}
			}
			return
		}

		for _, client := range clients {
			c.logger.Info(client.RemoteAddr().String())
			if err = client.WriteMessage(msgType, msg); err != nil {
				c.logger.Error(err)
				return
			}

		}

	}
}
