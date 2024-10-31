package tests

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			break
		}
		fmt.Printf("Received message: %s\n", message) // Print received message to console
		err = c.WriteMessage(mt, message)
		if err != nil {
			break
		}
	}
}

func TestExample(t *testing.T) {
	// Create test server with the echo handler.
	s := httptest.NewServer(http.HandlerFunc(echo))
	defer s.Close()

	// Convert http://127.0.0.1 to ws://127.0.0.
	u := "ws" + strings.TrimPrefix(s.URL, "http")

	// Connect to the server
	ws, _, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer ws.Close()

	// Send "hello world!" message to server, read response and check to see if it's what we expect.
	for i := 0; i < 10; i++ {
		message := "hello world!"
		if err := ws.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
			t.Fatalf("%v", err)
		}
		_, p, err := ws.ReadMessage()
		if err != nil {
			t.Fatalf("%v", err)
		}
		if string(p) != message {
			t.Fatalf("bad message")
		}
	}
}
