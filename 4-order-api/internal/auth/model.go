package auth

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Phone     string `json:"phone" validate:"e164, required" gorm:"uniqueIndex"`
	SessionID string `json:"sessionid"`
	Code      string `json:"code"`
}

func NewUser(phone, sessionID string, code string) *User {
	return &User{Phone: phone, SessionID: sessionID, Code: code}
}
