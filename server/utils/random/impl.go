package random

import "math/rand"

var alphaNum = "abcdefghijkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ23456789"

// generates a random string of fixed size
func RandomString(size int) string {
	buf := make([]byte, size)
	for i := 0; i < size; i++ {
		buf[i] = alphaNum[rand.Intn(len(alphaNum))]
	}
	return string(buf)
}
