package switchboard

import (
	"github.com/adamkrieger/symphony/common"
	"github.com/adamkrieger/symphony/quickconf/caller"
	"github.com/adamkrieger/symphony/quickconf/confcall"
	"github.com/adamkrieger/symphony/quickconf/contracts"
	"github.com/gorilla/websocket"
	"log"
)

const (
	switchBoardIDLength = 6
)

type switchboard struct {
	switchboardID  string
	confcalls      map[string]contracts.ConfCall
	newCallerQueue chan *caller.NewCaller
}

func NewSwitchboard() contracts.SwitchBoard {
	callMap := make(map[string]contracts.ConfCall)

	retSB := &switchboard{
		switchboardID:  string(common.RandASCIIBytes(switchBoardIDLength)),
		confcalls:      callMap,
		newCallerQueue: make(chan *caller.NewCaller, 50),
	}

	go retSB.processNewSocketConnections()

	return retSB
}

func (swBoard *switchboard) HandleNewSocketConnection(conn *websocket.Conn) {

	log.Println("handling new connection")

	newCaller, err := caller.ProcessNewCaller(conn)

	if err != nil {
		log.Println("Error processing new caller")
		return
	}

	swBoard.newCallerQueue <- newCaller
}

func (swBoard *switchboard) processNewSocketConnections() {
	for {
		select {
		case newCaller := <-swBoard.newCallerQueue:
			desiredCall, callExists := swBoard.confcalls[newCaller.DesiredCallID]

			if !callExists {
				desiredCall = confcall.NewConfCall(newCaller.DesiredCallID)
				swBoard.confcalls[newCaller.DesiredCallID] = desiredCall
			}

			desiredCall.AddToCall(newCaller.CallerID, newCaller.Conn)
		}
	}
}
