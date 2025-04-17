package database

import (
	"fmt"
	"log"
	"net/http"
	"real-time-forum/variables"
	"time"
)

func InsertUser(user *variables.User) {

	InsertData :=
		`
	INSERT INTO users
	VALUES (?, ?, ?, ?, ?, ?, ?, ?);
	`

	_, err := DB.Exec(InsertData, user.ID, user.Nickname, user.Age, user.Gender, user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		log.Fatal(err)
	}
}

func GetUserByEmail(email string) *variables.User {

	var user variables.User

	GetData :=
		`
	SELECT * FROM users
	WHERE email = ?;
	`
	Rows, err := DB.Query(GetData, email)
	if err != nil {
		log.Fatal(err)
	}
	for Rows.Next() {
		err = Rows.Scan(&user.ID, &user.Nickname, &user.Age, &user.Gender, &user.FirstName, &user.LastName, &user.Email, &user.Password)
		if err != nil {
			log.Fatal(err)
		}
	}
	return &user
}

func GetUserByNickname(nickname string) *variables.User {

	var user variables.User

	GetData :=
		`
	SELECT * FROM users
	WHERE nickname = ?;
	`
	Rows, err := DB.Query(GetData, nickname)
	if err != nil {
		log.Fatal(err)
	}
	for Rows.Next() {
		err = Rows.Scan(&user.ID, &user.Nickname, &user.Age, &user.Gender, &user.FirstName, &user.LastName, &user.Email, &user.Password)
		if err != nil {
			log.Fatal(err)
		}
	}
	return &user
}

func InsertPost(post *variables.Post) {

	InsertData :=
		`
	INSERT INTO posts (
		created_at, 
		title, 
		content, 
		category, 
		user_id
		)
	VALUES (?, ?, ?, ?, ?);
	`

	_, err := DB.Exec(InsertData, post.Date, post.Title, post.Content, post.Category, post.User.ID)
	if err != nil {
		log.Fatal(err)
	}
}

func GetPostByID(id int) *variables.Post {
	var post variables.Post

	GetData :=
		`
	SELECT * FROM posts
	WHERE id = ?;
	`
	Rows, err := DB.Query(GetData, id)
	if err != nil {
		log.Fatal(err)
	}
	for Rows.Next() {
		var userID string
		err = Rows.Scan(&post.ID, &post.Title, &post.Content, &post.Category, &userID)
		if err != nil {
			log.Fatal(err)
		}
		post.User = GetUserByID(userID)
	}
	return &post
}

func GetCurrentUser(r *http.Request) *variables.User {
	cookie, err := r.Cookie("session")
	if err != nil {
		fmt.Println("Error getting cookie:", err)
		return nil
	}

	var userID string

	err = DB.QueryRow("SELECT user_id FROM sessions WHERE id = ? AND expiration > ?", cookie.Value, time.Now()).Scan(&userID)
	if err != nil {
		fmt.Println("Error getting user ID from session:", err)
		return nil
	}

	user := GetUserByID(userID)
	return user
	
}

func GetAllUsers() []*variables.User {
	rows, err := DB.Query("SELECT * FROM users")
	if err != nil {
		log.Println("Error getting users:", err)
		return nil
	}
	defer rows.Close()

	var users []*variables.User
	
	for rows.Next() {
		var user variables.User
		err := rows.Scan(&user.ID, &user.Nickname, &user.Age, &user.Gender, &user.FirstName, &user.LastName, &user.Email, &user.Password)
		if err != nil {
			log.Println("Erreur lors du scan d’un utilisateur :", err)
			continue
		}
		users = append(users, &user)
}
	if err := rows.Err(); err != nil {
		log.Println("Error iterating over users:", err)
		return nil
	}
	return users
}


func GetUserByID(id string) *variables.User {

	var user variables.User

	GetData :=
		`
	SELECT * FROM users
	WHERE id = ?;
	`
	Rows, err := DB.Query(GetData, id)
	if err != nil {
		log.Fatal(err)
	}
	for Rows.Next() {
		password := []byte{}
		err = Rows.Scan(&user.ID, &user.Nickname, &user.Age, &user.Gender, &user.FirstName, &user.LastName, &user.Email, &password)
		if err != nil {
			log.Fatal(err)
		}
	}
	return &user
}

func GetpostHome() []*variables.Post {
	var posts []*variables.Post

	GetData :=
		`
	SELECT * FROM posts
	ORDER BY created_at DESC;
	`
	Rows, err := DB.Query(GetData)
	if err != nil {
		log.Fatal(err)
	}
	for Rows.Next() {
		var post variables.Post
		var userID string
		err = Rows.Scan(&post.ID, &post.Title, &post.Content, &post.Category, &userID, &post.Date)
		if err != nil {
			log.Fatal(err)
		}
		post.User = GetUserByID(userID)
		post.Date = time.Now()
		posts = append(posts, &post)
	}
	return posts
}

func InsertSession(session_token string, user *variables.User) {
	InsertData :=
		`
	INSERT INTO sessions
	VALUES (?, ?, ?);
	`

	expiration := time.Now().Add(time.Hour * 1)
	_, err := DB.Exec(InsertData, session_token, user.ID, expiration)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Session inserted")
	
}

func DeleteSession(session_token string) {
	DeleteData :=
		`
	DELETE FROM sessions
	WHERE id = ?;
	`

	_, err := DB.Exec(DeleteData, session_token)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Session deleted")
	
}