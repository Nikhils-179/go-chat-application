url = "ws://localhost:8080/ws"
let ws = new WebSocket(url)

ws.onmessage = function(event) {
    let data = JSON.parse(event.data);
    let msgDiv = document.createElement("div");
    msgDiv.className = "message received";
    
    if (data.history) {
        msgDiv.innerText = "History: " + data.history.join(", ");
    } else {
        msgDiv.innerText = "Received: " + data.message;
    }
    
    document.querySelector(".messages").appendChild(msgDiv);
    console.log("Received: ", data);
}

document.getElementById("sendButton").addEventListener("click", function() {
    sendMessage();
})

document.getElementById("messageInput").addEventListener("keydown", function(event) {
    if (event.key === "Enter") {
        sendMessage();
    }
})

document.getElementById("getHistoryButton").addEventListener("click", function() {
    getHistory();
})

function sendMessage() {
    let input = document.getElementById("messageInput");
    let message = input.value;
    ws.send(message);

    let msgDiv = document.createElement("div");
    msgDiv.className = "message sent";
    msgDiv.innerText = "Sent: " + message;
    document.querySelector(".messages").appendChild(msgDiv);

    input.value = "";
}

function getHistory() {
    ws.send("GET_HISTORY");
}
