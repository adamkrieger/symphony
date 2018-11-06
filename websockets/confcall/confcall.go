package confcall

import (
	"github.com/adamkrieger/symphony/common"
	"github.com/adamkrieger/symphony/websockets/caller"
	"github.com/adamkrieger/symphony/websockets/contracts"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

type confCall struct {
	callID        string
	callers       map[string]caller.Caller
	broadcastChan chan string
	newCallerChan chan caller.Caller
}

func NewConfCall() contracts.ConfCall {
	retConfCall := &confCall{
		callID:        string(common.RandASCIIBytes(6)),
		callers:       make(map[string]caller.Caller, 50),
		broadcastChan: make(chan string, 50),
		newCallerChan: make(chan caller.Caller, 50),
	}

	go retConfCall.startRouting()
	go retConfCall.printStatusRepeatedly()

	return retConfCall
}

func (conference *confCall) AddToCall(conn *websocket.Conn) {
	sessionID := string(common.RandASCIIBytes(6))

	newCaller := caller.ConnectCaller(sessionID, conn)

	conference.newCallerChan <- newCaller
}

func (conference *confCall) startRouting() {
	for {
		select {
		case newCaller := <-conference.newCallerChan:
			conference.callers[newCaller.SessionID()] = newCaller
			go conference.listenToNewCaller(newCaller)

			welcomeMsg := "welcome to callID " + conference.callID + ", your sessionID is " + newCaller.SessionID()
			newCaller.ToCaller() <- welcomeMsg

			log.Println("caller added: ", newCaller.SessionID())

		case broadcastMsg := <-conference.broadcastChan:
			for sessionID, eachCaller := range conference.callers {
				if !eachCaller.Disconnecting() {
					eachCaller.ToCaller() <- broadcastMsg
					log.Printf("msg send to %s: %s", sessionID, broadcastMsg)
				} else {
					delete(conference.callers, sessionID)

					log.Println("caller removed: ", sessionID)
				}
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

func (conference *confCall) printStatusRepeatedly() {
	for {
		msg := time.Now().String() + " PING"
		conference.broadcastChan <- msg

		time.Sleep(1 * time.Second)
	}
}
