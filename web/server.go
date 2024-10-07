package web

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gorilla/websocket"
	"gorm.io/gorm"
)

// Upgrades an HTTP server connection to the WebSocket protocol.
// After that, the HTTP connection is a WebSocket connection.
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func ServeHTTP(db *gorm.DB, port int) error {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		conn, err := upgrader.Upgrade(writer, request, nil)
		if err != nil {
			slog.Warn("Failed to upgrade connection", "addr", request.RemoteAddr)
			return
		}

		defer conn.Close()
		for HandleWebsocketConn(db, conn) == nil {
		}
	})

	slog.Debug("Serving HTTP", "port", port)
	return http.ListenAndServe(fmt.Sprintf("127.0.0.1:%d", port), nil)
}
