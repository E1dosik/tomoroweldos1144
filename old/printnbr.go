package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	if n > 0 {
		n = -n
	} else {
		z01.PrintRune('-')
	}
	d := 1
	for n/d <= -10 {
		d *= 10
	}
	for d > 0 {
		digit := -(n / d % 10)
		r := '0'
		for i := 0; i < digit; i++ {
			r++
		}
		z01.PrintRune(r)
		d /= 10
	}
}
