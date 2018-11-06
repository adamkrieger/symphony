package caller

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

type Caller interface {
	SessionID() string
	ToCaller() chan<- string
	FromCaller() <-chan string
	Disconnect()
	Disconnecting() bool
}

type caller struct {
	sessionID string
	socket           *websocket.Conn
	toSocketInput    chan<- string
	toSocketOutput   <-chan string
	fromSocketInput  chan<- string
	fromSocketOutput <-chan string
	sessionStart     time.Time
	isDisconnecting  bool
}

func ConnectCaller(sessionID string, conn *websocket.Conn) Caller {
	toSocketChan := make(chan string)
	fromSocketChan := make(chan string)

	retCaller := &caller{
		sessionID:        sessionID,
		socket:           conn,
		sessionStart:     time.Now().UTC(),
		toSocketInput:    toSocketChan,
		toSocketOutput:   toSocketChan,
		fromSocketInput:  fromSocketChan,
		fromSocketOutput: fromSocketChan,
		isDisconnecting:  false,
	}

	go retCaller.listenToConnection()
	go retCaller.listenToQueue()

	return retCaller
}

func (caller *caller) SessionID() string {
	return caller.sessionID
}

func (caller *caller) ToCaller() chan<- string {
	return caller.toSocketInput
}

func (caller *caller) FromCaller() <-chan string {
	return caller.fromSocketOutput
}

func (caller *caller) Disconnect() {
	caller.isDisconnecting = true
}

func (caller *caller) Disconnecting() bool {
	return caller.isDisconnecting
}

func (caller *caller) listenToConnection() {
	defer caller.socket.Close()
	defer close(caller.fromSocketInput)

	for {
		msg, readErr := readOneFromSocket(caller.socket)

		if readErr != nil {
			log.Println("Socket Read Err")
			caller.Disconnect()
			return
		} else {
			caller.fromSocketInput <- string(msg)
		}
	}
}

func (caller *caller) listenToQueue() {
	for {
		msg, queueOK := <-caller.toSocketOutput

		if !queueOK {
			log.Println("Caller Queue Closed")
			return
		} else {
			writeOneToSocket(caller.socket, []byte(msg))
		}
	}
}

func readOneFromSocket(socket *websocket.Conn) ([]byte, error) {
	_, message, err := socket.ReadMessage()
	return message, err
}

func writeOneToSocket(socket *websocket.Conn, message []byte) error {
	ourMap := map[string]string{"content": string(message)}
	jsonMessage, _ := json.Marshal(ourMap)

	return socket.WriteMessage(websocket.TextMessage, []byte(jsonMessage))
}
