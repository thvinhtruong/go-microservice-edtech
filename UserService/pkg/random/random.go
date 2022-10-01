package random

import (
	"math/rand"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func RandomName() string {
	return RandomString(7)
}

func RandomGender() string {
	genders := []string{"male", "female", "other"}
	return genders[rand.Intn(2)]
}

func RandomEmail() string {
	return RandomString(10) + "@example.com"
}

func RandomPhone() string {
	return strconv.Itoa(int(RandomInt(1000000000, 9999999999)))
}

func RandomDate() time.Time {
	return time.Unix(RandomInt(0, time.Now().Unix()), 0)
}
