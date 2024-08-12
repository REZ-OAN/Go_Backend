package utils

import (
	"math/rand"
	"strings"
	"time"
)

const alphabets = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomly generates an integer number between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// randomly generates a string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabets)

	for i := 0; i < n; i++ {
		c := alphabets[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()

}

// random Owner generate
func RandomOwner() string {
	return RandomString(6)
}

// random money generate
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// random currency generation
func RandomCurrency() string {
	currencies := []string{"BDT", "USD"}

	n := len(currencies)
	return currencies[rand.Intn(n)]
}

// random amount generation including negative number
func RandomAmount() int64 {
	return RandomInt(-2000, 2000)
}
