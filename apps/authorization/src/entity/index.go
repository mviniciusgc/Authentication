package entity

type UserResponse struct {
	ID string `json:"id,omitempty"`
}
type TokenResponse struct {
	AccessToken      string `json:"access_token,omitempty"`
	IdToken          string `json:"id_token,omitempty"`
	ExpiresIn        int    `json:"expires_in,omitempty"`
	RefreshExpiresIn int    `json:"refresh_expires_in,omitempty"`
	RefreshToken     string `json:"refresh_token,omitempty"`
	TokenType        string `json:"token_type,omitempty"`
	NotBeforePolicy  int    `json:"not-before-policy,omitempty"`
	SessionState     string `json:"session_state,omitempty"`
	Scope            string `json:"scope,omitempty"`
}

type AuthenticateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}
type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}
type UserUpdateRequest struct {
	Email *string `json:"email"`
}
