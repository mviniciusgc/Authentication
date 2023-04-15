package entity

type UserResponse struct {
	ID string `json:"id,omitempty"`
}
type TokenResponse struct {
	AccessToken      string `json:"access_token"`
	IdToken          string `json:"id_token"`
	ExpiresIn        int    `json:"expires_in"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
	RefreshToken     string `json:"refresh_token"`
	TokenType        string `json:"token_type"`
	NotBeforePolicy  int    `json:"not-before-policy"`
	SessionState     string `json:"session_state"`
	Scope            string `json:"scope"`
}

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}
type AuthenticateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
