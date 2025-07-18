package entities

import "time"

type Otp struct {
	Email     string
	Code      int
	ExpiresAt time.Time
}

func NewOtp(email string, code int, expiresAt time.Time) *Otp {
	return &Otp{email, code, expiresAt}
}
