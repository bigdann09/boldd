package entities

import (
	"time"

	"github.com/google/uuid"
)

type Otp struct {
	UUID      string
	Email     string
	Code      int
	ExpiresAt time.Time
}

func NewOtp(email string, code int, expiresAt time.Time) *Otp {
	return &Otp{
		UUID:      uuid.NewString(),
		Email:     email,
		Code:      code,
		ExpiresAt: expiresAt,
	}
}
