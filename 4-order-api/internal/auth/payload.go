// Package auth
package auth

type AuthRequest struct {
	Phone string `json:"phone" validate:"e164"`
}

type AuthPreResponse struct {
	SessionID string `json:"sessionid"`
}

type AuthResponse struct {
	AuthPreResponse
	Code string `json:"code"`
}

func NewAuthRequest(phone string) *AuthRequest {
	return &AuthRequest{Phone: phone}
}

func NewAuthPreResponse(s string) *AuthPreResponse {
	return &AuthPreResponse{SessionID: s}
}

func NewAuthResponse(r AuthPreResponse) *AuthResponse {
	return &AuthResponse{r, "3456"}
}
