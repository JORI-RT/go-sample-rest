package utsample

import "strconv"

// ３で割り切れるならfizz、５で割り切れるならbuzz、１5で割り切れるならfizzbussを返す
func fizzBuzz(i int) string {
	if i%15 == 0 {
		return "fizzBuzz"
	} else if i%5 == 0 {
		return "buzz"
	} else if i%3 == 0 {
		return "fizz"
	}
	return strconv.Itoa(i)
}
