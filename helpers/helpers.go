package helpers

import "math/rand"

type SomeType struct {
	TypeName string
}

func RandomNumber(n int) int {
	value := rand.Intn(n)
	return value
}
