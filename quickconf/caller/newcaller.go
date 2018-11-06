package caller

import "github.com/gorilla/websocket"

type NewCaller struct {
	CallerID      string
	DesiredCallID string
	Conn          *websocket.Conn
}

func ProcessNewCaller(conn *websocket.Conn) (*NewCaller, error) {
	callerID, callerIDErr := readOneFromSocket(conn)

	if callerIDErr != nil {
		conn.Close()
		return nil, callerIDErr
	}

	desiredCall, desiredCallErr := readOneFromSocket(conn)

	if desiredCallErr != nil {
		conn.Close()
		return nil, desiredCallErr
	}

	return &NewCaller{
		Conn:          conn,
		CallerID:      string(callerID),
		DesiredCallID: string(desiredCall),
	}, nil
}
