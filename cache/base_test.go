package cache_test

import "math/rand"

const testFailedMsg string = "%s failed; want %v but got %v"

func randint(atleast, atmost int) int {
	return rand.Intn(atmost-atleast) + atleast
}
