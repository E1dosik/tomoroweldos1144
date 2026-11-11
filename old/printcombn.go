package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	var c [10]rune
	r := '0'
	for i := 0; i < n; i++ {
		c[i] = r
		r++
	}
	c[n] = ':'
	for {
		for i := 0; i < n; i++ {
			z01.PrintRune(c[i])
		}
		i := n - 1
		for i >= 0 && c[i]+1 >= c[i+1] {
			i--
		}
		if i < 0 {
			break
		}
		c[i]++
		for j := i + 1; j < n; j++ {
			c[j] = c[j-1] + 1
		}
		z01.PrintRune(',')
		z01.PrintRune(' ')
	}
	z01.PrintRune('\n')
}
