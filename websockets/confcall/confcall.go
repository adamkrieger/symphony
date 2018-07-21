package confcall

import (
	"github.com/gorilla/websocket"
	"github.com/adamkrieger/symphony/websockets/contracts"
	"github.com/adamkrieger/symphony/websockets/caller"
	"github.com/adamkrieger/symphony/common"
	"log"
)

type confCall struct {
	callers map[string]caller.Caller
	broadcastChan chan string
	newCallerChan chan caller.Caller
}

func NewConfCall() contracts.ConfCall {
	retConfCall := &confCall{
		callers:       make(map[string]caller.Caller, 50),
		broadcastChan: make(chan string),
		newCallerChan: make(chan caller.Caller),
	}

	go retConfCall.startRouting()

	return retConfCall
}

func (conference *confCall) AddToCall(conn *websocket.Conn) {
	sessionID := string(common.RandASCIIBytes(6))

	newCaller := caller.ConnectCaller(sessionID, conn)

	conference.newCallerChan <- newCaller
}

func (conference *confCall) startRouting() {
	select {
	case newCaller := <-conference.newCallerChan:
		conference.callers[newCaller.SessionID()] = newCaller
		go conference.listenToNewCaller(newCaller)

		log.Println("caller ", newCaller.SessionID(), " added")

	case broadcastMsg := <-conference.broadcastChan:
		for sessionID, eachCaller := range conference.callers {
			if !eachCaller.Disconnecting() {
				eachCaller.ToCaller() <- broadcastMsg
			} else {
				delete(conference.callers, sessionID)
			}
		}
	}
}

func (conference *confCall) listenToNewCaller(newCaller caller.Caller) {
	for {
		select {
		case msg, chanOK := <-newCaller.FromCaller():
			if chanOK {
				conference.broadcastChan <- msg
			} else {
				return
			}
		}
	}
}