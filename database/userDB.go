package database

import (
	"log"
	"real-time-forum/variables"
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

