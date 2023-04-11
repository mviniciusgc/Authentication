package entity

type UserResponse struct {
	ID string `json:"id,omitempty"`
}

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"user_name"`
}
