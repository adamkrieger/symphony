package contracts

import "github.com/gorilla/websocket"

type SwitchBoard interface {
	HandleNewSocketConnection(conn *websocket.Conn)
}

type ConfCall interface {
	AddToCall(callerID string, conn *websocket.Conn)
}

type NewCallHandler func(conn *websocket.Conn)
