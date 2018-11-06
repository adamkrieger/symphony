package switchboard

import (
	"github.com/adamkrieger/symphony/common"
	"github.com/adamkrieger/symphony/websockets/confcall"
	"github.com/adamkrieger/symphony/websockets/contracts"
	"github.com/gorilla/websocket"
	"log"
)

type switchboard struct {
	switchboardID string
	monocall      contracts.ConfCall
}

func NewSwitchboard() contracts.SwitchBoard {
	singleCall := confcall.NewConfCall()

	retSB := &switchboard{
		switchboardID: string(common.RandASCIIBytes(6)),
		monocall:      singleCall,
	}

	return retSB
}

func (swBoard *switchboard) HandleNewSocketConnection(conn *websocket.Conn) {

	log.Println("handling new connection")

	swBoard.monocall.AddToCall(conn)
}
