// export function loadHome() {
//   fetch("/home")
//     .then((response) => {
//       if (!response.ok) throw new Error("Erreur lors du chargement des posts");
//       console.log(response.json())
//       return response.json(); // car /home renvoie du HTML
//     })
//     .then((posts) => {
//       console.log("HTML reçu :", html);
//       const app = document.getElementById("app");
//       app.innerHTML = formatPosts(posts);

//     })
//     .catch((error) => {
//       console.error("Erreur :", error);
//       document.getElementById("app").innerHTML =
//         "<p>Impossible de charger les posts.</p>";
//     });
// }

export async function loadHome() {
  let resp = await fetch("/home");
  let r = await resp.json();
  let posts = r.Posts;
  const app = document.getElementById("app");
  app.innerHTML = formatPosts(posts);
}

function formatPosts(posts) {
  let result = "";
  for (let i = 0; i < posts.length; i++) {
    let post = posts[i];
    let postHTML = `
      <div class="post">
      <h1 class="title">${post.title}</h1>
      <h2 class="user">${post.user.nickname}</h2>
    
      <p class="content">${post.content}</p>

      <div class="footer">
        <span class="date">${post.date}</span>
        <span class="category">${post.category}</span>
      </div>
    </div>
    `;
    result += postHTML;
  }
  return result;
}
