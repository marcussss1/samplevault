package model

type SignupUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type FullUser struct {
	ID        string `json:"id"`
	SessionID string `json:"session_id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

type User struct {
	ID        string `json:"id"`
	SessionID string `json:"session_id"`
	Username  string `json:"username"`
}
