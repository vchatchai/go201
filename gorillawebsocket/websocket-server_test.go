package gorillawebsocket

import (
	"github.com/gorilla/websocket"
	"github.com/magiconair/properties/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestWebSocketServer(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test WebSocketServer"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WebSocketServer()
		})
	}
}

func TestWebSocketServerAuto(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(HandleClients))
	defer server.Close()

	u := "ws" + strings.TrimPrefix(server.URL, "http")
	socket, _, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer socket.Close()

	m := Message{Message: "hello"}
	if err := socket.WriteJSON(&m); err != nil {
		t.Fatalf("%v", err)
	}
	var message Message
	err = socket.ReadJSON(&message)
	if err != nil {
		t.Fatalf("%v", err)
	}

	assert.Equal(t, "hello", message.Message, "They should be equal")

}
