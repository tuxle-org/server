package tuxle

import (
	"github.com/bbfh-dev/go-tools/tools/terr"
	"github.com/gorilla/websocket"
	"github.com/tuxle-org/lib/tuxle/protocol"
	"gorm.io/gorm"
)

var OK protocol.Letter = nil

func Handle(db *gorm.DB, conn *websocket.Conn, letter protocol.Letter) protocol.Letter {
	terr.Assert(conn != nil, "Connection must exist")
	terr.Assert(letter != nil, "Letter must be a valid type, not nil")

	switch letter := letter.(type) {
	case protocol.GetLetter:
		switch letter.Query {
		case protocol.GET_SERVER_INFO:
			// TODO: Caching would be nice, especially knowing how often this query will probably be sent!
			serverInfo, err := ServerInfo(db)
			if err != nil {
				return protocol.ErrLetter{Body: err.Error()}
			}
			letter, err := protocol.EncodeEntity(serverInfo)
			if err != nil {
				return protocol.ErrLetter{Body: err.Error()}
			}
			return letter
		}
	}

	return OK
}
