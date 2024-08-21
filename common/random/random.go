package random

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"
const numbers = "0123456789"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomInt32(min, max int32) int32 {
	return int32(RandomInt(int64(min), int64(max)))
}

func RandomFloat(min, max int64) float32 {
	return float32(min + rand.Int63n(max-min+1))
}

func RandomString(n int) string {
	var sb strings.Builder

	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomPhone() string {
	var sb strings.Builder
	sb.WriteByte('0')
	sb.WriteByte('1')
	sb.WriteByte(numbers[rand.Intn(3)])

	k := len(numbers)
	for i := 0; i < 8; i++ {
		c := numbers[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomName() string {
	return RandomString(7)
}

func RandomMoney() float32 {
	return RandomFloat(3, 1000)
}

func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}

func RandomURL() string {
	return fmt.Sprintf("http://%s.com", RandomString(7))
}
