export function loadHome() {
  fetch("/home")
    .then((response) => {
      if (!response.ok) throw new Error("Erreur lors du chargement des posts");
      return response.text(); // car /home renvoie du HTML
    })
    .then((html) => {
      const app = document.getElementById("app");
      app.innerHTML = html;
    })
    .catch((error) => {
      console.error("Erreur :", error);
      document.getElementById("app").innerHTML =
        "<p>Impossible de charger les posts.</p>";
    });

  fetch("/home")
    .then((response) => {
      if (!response.ok)
        throw new Error("Erreur lors du chargement des Utilisateurs");
      return response.text(); // car /home renvoie du HTML
    })
    .then((html) => {
      const app = document.getElementById("app");
      app.innerHTML = html;
    })
    .catch((error) => {
      console.error("Erreur :", error);
      document.getElementById("app").innerHTML =
        "<p>Impossible de charger les Utilisateurs.</p>";
    });
}
