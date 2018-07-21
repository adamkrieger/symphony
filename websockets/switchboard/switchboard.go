package switchboard

import (
	"github.com/adamkrieger/symphony/websockets/contracts"
	"github.com/adamkrieger/symphony/websockets/confcall"
	"github.com/gorilla/websocket"
	"log"
)

type switchboard struct {
	monocall contracts.ConfCall
}

func NewSwitchboard() contracts.SwitchBoard {
	singleCall := confcall.NewConfCall()

	retSB := &switchboard{
		monocall:singleCall,
	}

	return retSB
}

func (swBoard *switchboard) HandleNewSocketConnection(conn *websocket.Conn) {

	log.Println("handling new connection")

	swBoard.monocall.AddToCall(conn)
}