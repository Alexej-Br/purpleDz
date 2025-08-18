package auth

type LoginResponse struct {
	Token string `json:"token"`
}

type VerifyResponse struct {
	Hash string `json:"hash"`
}
