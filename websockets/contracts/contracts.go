package contracts

import "github.com/gorilla/websocket"

type SwitchBoard interface {
	HandleNewSocketConnection(conn *websocket.Conn)
}

type ConfCall interface {
	AddToCall(conn *websocket.Conn)
}
