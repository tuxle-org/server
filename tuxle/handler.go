package tuxle

import (
	"github.com/bbfh-dev/go-tools/tools/terr"
	"github.com/gorilla/websocket"
	"github.com/tuxle-org/lib/tuxle/protocol"
)

var OK protocol.Letter = nil

func Handle(conn *websocket.Conn, letter protocol.Letter) protocol.Letter {
	terr.Assert(conn != nil, "Connection must exist")
	terr.Assert(letter != nil, "Letter must be a valid type, not nil")

	switch letter.(type) {
	}

	return OK
}
