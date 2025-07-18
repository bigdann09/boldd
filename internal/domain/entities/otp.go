package entities

import "time"

type Otp struct {
	UserID    uint
	Code      int
	ExpiresAt time.Time
}

func NewOtp(userID uint, code int, expiresAt time.Time) *Otp {
	return &Otp{userID, code, expiresAt}
}
