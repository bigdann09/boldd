package utils

import "math/rand"

func GenerateOTP() int {
	return rand.Intn(999999)
}
