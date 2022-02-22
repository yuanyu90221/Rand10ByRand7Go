package random_util

import "math/rand"

func Rand10() int {
	for {
		a := rand7()
		b := rand7()
		val := b + (a-1)*7
		if val <= 40 {
			mod := val % 10
			if mod == 0 {
				return 10
			}
			return mod
		}
	}
}

func rand7() int {
	return rand.Intn(7) + 1
}
