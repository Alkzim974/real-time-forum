export function InitWS() {
  const ws = new WebSocket("ws://localhost:8080/ws");
  ws.onopen = function (event) {
    console.log("WebSocket is open now.");
  };
  ws.onclose = function (event) {
    console.log("WebSocket is closed now.");
    if (event.wasClean) {
      console.log("Connection closed cleanly");
    }
  };
  ws.onmessage = function (event) {
    console.log("Message from server ", event.data);

    const data = JSON.parse(event.data);
    if (data.type === "message") {
      const message = data.message;
      const chatBox = document.getElementById("chat-box");
      chatBox.innerHTML += `<p>${message}</p>`;
    }
  };
}
