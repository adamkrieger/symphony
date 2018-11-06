Symphony
--------

#### What is Symphony?

Symphony is a set of samples experimenting with service scaling in distributed architecture.

#### What do I do first?

Prerequisites

-  Docker

Steps

1. `/traefik/run-it.sh` in a terminal
   - `localhost:8080` becomes traefik admin portal
2. `/restapi/compose-up.sh` in a terminal
   - `http://api.docker.localhost:8081` becomes sample rest api
   - Test with `client.html - Balanced Rest API - Send`
   - Test a few times to watch the balancer cycle
3. `/websockets/compose-up.sh` in a terminal
   - `ws://confcall.docker.localhost:8082` becomes confcall socket
   - Test with `client.html - ConfCall Dialer - Open Socket / Send`
   - *Currently Connects to Random Call*