package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	fmt.Printf("%d\n", commonMultiple(a, b))
}

func getDif(small, bigger int) int {
	s := small
	b := bigger
	if s > b {
		b = small
		s = bigger
	}

	dif := b - s
	if s == dif {
		return dif
	}
	return getDif(s, dif)
}

func commonDivisor(x, y int) int {
	if x%2 == 0 && y%2 == 0 {
		return 2 * commonDivisor(x/2, y/2)
	}
	return getDif(x, y)
}

func commonMultiple(x, y int) int {
	return x * y / commonDivisor(x, y)
}
