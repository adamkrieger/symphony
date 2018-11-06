package main

import (
	"github.com/adamkrieger/symphony/websockets/contracts"
	"github.com/adamkrieger/symphony/websockets/switchboard"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Println("Starting Switchboard")

	mainSwitchboard := switchboard.NewSwitchboard()

	go launchSocketListener(mainSwitchboard)

	waitChan := make(chan interface{})
	<-waitChan
}

func launchSocketListener(swboard contracts.SwitchBoard) {
	socketUpgrader := &websocket.Upgrader{
		CheckOrigin:     alwaysAccept,
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	socketHandler := buildSocketHandler(swboard, socketUpgrader)

	http.HandleFunc("/ws", socketHandler)
	port := ":" + os.Getenv("RUNPORT")

	log.Println("Launching Socket Handler at ", port)

	serveErr := http.ListenAndServe(port, nil)
	if serveErr != nil {
		os.Exit(1)
	}
}

func buildSocketHandler(swboard contracts.SwitchBoard, upgrader *websocket.Upgrader) func(w http.ResponseWriter, r *http.Request) {
	return func(respWriter http.ResponseWriter, req *http.Request) {
		socket, upgradeErr := upgrader.Upgrade(respWriter, req, nil)
		if upgradeErr != nil {
			http.NotFound(respWriter, req)
		}

		swboard.HandleNewSocketConnection(socket)
	}
}

func alwaysAccept(req *http.Request) bool { return true }
