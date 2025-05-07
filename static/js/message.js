import { sendPrivateMessage, ws } from "./websocket.js";

export async function chatBox(nickname) {
  const app = document.getElementById("users");
  app.innerHTML = `
      <h2>Chat avec ${nickname}</h2>
      <div id="chatBox"></div>
      <input type="text" id="messageInput" placeholder="Votre message...">
      <button id="sendMessage">Envoyer</button>
    `;
  const sendMessageButton = document.getElementById("sendMessage");
  const messageInput = document.getElementById("messageInput");
  const chatBox = document.getElementById("chatBox");
  chatBox.dataset.nickname = nickname;

  sendMessageButton.addEventListener("click", () => {
    const message = messageInput.value;
    if (message) {
      sendPrivateMessage(nickname, message);
      chatBox.innerHTML += `<p><strong>Vous:</strong> ${message}</p>`;
      messageInput.value = ""; // Clear the input after sending
    }
  });

  fetchMessages(nickname)
    .then((messages) => {
      messages.forEach((message) => {
        chatBox.innerHTML += `<p><strong>${message.sender}:</strong> ${message.content}</p>`;
      });
    })
    .catch((error) => {
      console.error("Erreur lors de la récupération des messages", error);
    });
}

async function fetchMessages(nickname) {
  const response = await fetch(`/messages/${nickname}`);
  if (response.ok) {
    const messages = await response.json();
    return messages;
  } else {
    console.error("Erreur lors de la récupération des messages");
    return [];
  }
}
