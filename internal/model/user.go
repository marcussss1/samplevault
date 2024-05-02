package model

type SignupUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignupUserResponse struct {
	ID        string `json:"id"`
	SessionID string `json:"session_id"`
	Username  string `json:"username"`
}

type LoginUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginUserResponse struct {
	ID        string `json:"id"`
	SessionID string `json:"session_id"`
	Username  string `json:"username"`
}

type FullUser struct {
	ID        string `json:"id"`
	SessionID string `json:"session_id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

type GetUserBySessionID struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}
