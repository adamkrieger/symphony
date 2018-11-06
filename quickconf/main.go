package main

import (
	"github.com/adamkrieger/symphony/quickconf/contracts"
	"github.com/adamkrieger/symphony/quickconf/switchboard"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Println("Starting Switchboard")

	mainSwitchboard := switchboard.NewSwitchboard()

	go launchSocketListener(mainSwitchboard.HandleNewSocketConnection)

	waitChan := make(chan interface{})
	<-waitChan
}

func launchSocketListener(newSocketHandler contracts.NewCallHandler) {
	socketUpgrader := &websocket.Upgrader{
		CheckOrigin:     alwaysAccept,
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	wsEndpointHandler := buildWSEndpointHandler(newSocketHandler, socketUpgrader)

	http.HandleFunc("/ws", wsEndpointHandler)
	port := ":" + os.Getenv("RUNPORT")

	log.Println("Launching Socket Handler at ", port)

	serveErr := http.ListenAndServe(port, nil)
	if serveErr != nil {
		os.Exit(1)
	}
}

func buildWSEndpointHandler(newSocketHandler contracts.NewCallHandler, upgrader *websocket.Upgrader) func(w http.ResponseWriter, r *http.Request) {
	return func(respWriter http.ResponseWriter, req *http.Request) {
		socket, upgradeErr := upgrader.Upgrade(respWriter, req, nil)
		if upgradeErr != nil {
			http.NotFound(respWriter, req)
		}

		newSocketHandler(socket)
	}
}

func alwaysAccept(req *http.Request) bool { return true }
