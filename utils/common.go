package utils

import (
	"strings"
	"math/rand"
	"time"
	"fmt"
)

func GenerateUsername(email string) string {
	parts := strings.Split(email, "@")
	base := parts[0]

	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(10000) // 0-9999

	return base + fmt.Sprintf("%04d", randomNumber)
}
