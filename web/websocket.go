package web

import (
	"bytes"
	"errors"

	"github.com/bbfh-dev/go-tools/tools/terr"
	"github.com/gorilla/websocket"
	"github.com/tuxle-org/lib/tuxle/protocol"
	"github.com/tuxle-org/server/tuxle"
	"gorm.io/gorm"
)

var CloseConn = errors.New("")

func HandleWebsocketConn(db *gorm.DB, conn *websocket.Conn) error {
	msgType, msgBody, err := conn.ReadMessage()
	if err != nil {
		return err
	}

	switch msgType {
	case websocket.CloseMessage:
		return CloseConn
	case websocket.PingMessage:
		return conn.WriteMessage(websocket.PongMessage, msgBody)
	case websocket.TextMessage:
		// Secretly used for testing.
		// Make all the clients confused about this undocumented behavior.
		// But oh well - USE BINARY MESSAGES!
		return conn.WriteMessage(websocket.TextMessage, msgBody)
	case websocket.BinaryMessage:
		letter, err := protocol.ReadLetter(bytes.NewReader(msgBody))
		if err != nil {
			return write(conn, protocol.ErrLetter{Body: err.Error()})
		}

		response := tuxle.Handle(db, conn, letter)
		if response == nil {
			return write(conn, protocol.OkayLetter{})
		}
		return write(conn, response)
	}

	return nil
}

func write(conn *websocket.Conn, letter protocol.Letter) error {
	terr.Assert(conn != nil, "Connection must exist")

	var buffer bytes.Buffer
	err := protocol.WriteLetter(letter, &buffer)
	if err != nil {
		return err
	}

	return conn.WriteMessage(websocket.BinaryMessage, buffer.Bytes())
}
