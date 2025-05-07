import { displayUsers } from "./home.js";

export let ws = null;

export function InitWS() {
  ws = new WebSocket("ws://localhost:8080/ws");

  ws.onopen = function (event) {
    console.log("WebSocket is open now.");
  };

  ws.onclose = function (event) {
    console.log("WebSocket is closed now.");
    console.log("Message from server ", event);
  };

  ws.onmessage = function (event) {
    console.log("Message from server ", JSON.parse(event.data));

    try {
      const data = JSON.parse(event.data);
      // Si le message contient une information de connexion/déconnexion
      if (data.type === "log") {
        console.log("User connection update:", data.connexion);
        // Mettre à jour immédiatement la liste des utilisateurs
        displayUsers();
      }
      // Si le message contient une information de message
      else if (data.type === "message") {
        const chatBox = document.getElementById("chatBox");
        if (chatBox && chatBox.dataset.nickname === data.sender) {
          chatBox.innerHTML += `<p><strong>${data.sender}:</strong> ${data.content}</p>`;
        }
      }
    } catch (e) {
      console.error("Error parsing message:", e);
    }
  };

  ws.onerror = (error) => {
    console.error("WebSocket error observed:", error);
    ws.close();
  };
}

export function sendPrivateMessage(nickname, message) {
  console.log(ws);
  if (ws && ws.readyState === WebSocket.OPEN) {
    ws.send(
      JSON.stringify({
        type: "message",
        receiver: nickname,
        content: message,
        created_at: new Date(),
      })
    );
  } else {
    console.error("WebSocket is not open. Cannot send message.");
  }
}
