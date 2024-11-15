package rand

import "math/rand"

func Int(max int) int {
	return rand.Intn(max)
}
