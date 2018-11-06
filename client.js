let wsRef;

// function wireHandlers(){
//     let input = document.getElementById("message");
//
//     input.addEventListener("keyup", function(event) {
//         event.preventDefault();
//         if (event.key === 'Enter') {
//             document.getElementById("send").click();
//         }
//     });
// }

function WebSocketConnect() {
    if("WebSocket" in window) {
        //let ws = new WebSocket("ws://localhost:8053/ws");
        let ws = new WebSocket("ws://confcall.docker.localhost:8082/ws");

        ws.onopen = function(){
            console.log("Websocket open.");
            ws.send("hello from client");
            wsRef = ws;
        };

        ws.onmessage = function(msg) {
            let msgData = msg.data;
            console.log("Message Received: " + msgData)
        };

        ws.onclose = function(){
            console.log("WebSocket closed.")
        };

    } else {
        console.log("Browser does not support sockets.")
    }
}

function SendMessage(){
    let msgBox = document.getElementById("message");
    wsRef.send(msgBox.value)
}

function RestAPI_GET(){
    let req = new XMLHttpRequest();
    let url = document.getElementById("restapiaddr").value;
    console.log("REST GET @ ", url);
    req.open("GET", url, true);
    req.onload = function (e) {
        if (req.readyState === 4 && req.status === 200) {
            document.getElementById("restapi_resp").value = req.responseText;
        } else {
            console.log(req.statusText);
        }
    };
    req.send(null)
}