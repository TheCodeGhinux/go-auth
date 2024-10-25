package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// GenerateUsername generates a username based on the user's first name and a random number.
func GenerateUsername(firstName string) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNumber := r.Intn(1000)
	username := fmt.Sprintf("%s%d", firstName, randomNumber)
	return username
}
