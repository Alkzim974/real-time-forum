package variables

import "time"

type User struct {
	ID        string `json:"id"`
	Nickname  string `json:"nickname"`
	Age       int    `json:"age"`
	Gender    string `json:"gender"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  []byte `json:"password"`
}

type Post struct {
	ID       int
	Title    string
	Content  string
	Category string
	Date 	 time.Time
	User   	 *User // UUID
}

type UserStatus struct {
	Nickname       string `json:"nickname"`
	Online         bool   `json:"online"`
	HasNewMessage  bool   `json:"hasNewMessage"`
}
