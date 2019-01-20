package gorillawebsocket

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)
var upgrader = websocket.Upgrader{}

type Message struct {
	Message string `json:"message"`
}

func HandleClients(w http.ResponseWriter, r *http.Request) {
	go broadcastMessagesToClients()
	websocket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("error upgrading GET request to a websocket ::", err)
	}
	defer websocket.Close()

	clients[websocket] = true
	for {
		var message Message
		err := websocket.ReadJSON(&message)
		if err != nil {
			log.Printf("errors occurred while reading message: %v", err)
			delete(clients, websocket)
			break
		}
		broadcast <- message
	}
}

func broadcastMessagesToClients() {
	for {
		message := <-broadcast
		for client := range clients {
			err := client.WriteJSON(message)
			if err != nil {
				log.Printf("error occured while writing message to client: %v", err)

				client.Close()
				delete(clients, client)
			}
		}
	}
}

func WebSocketServer() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, "index.html")
	})

	http.HandleFunc("/echo", HandleClients)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("eror starting http server :: ", err)
		return
	}
}
