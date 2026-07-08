package main

import (
	"golang.org/x/tour/pic"
	"math"
)

func Pic(dx, dy int) [][]uint8 {
	s := make([][]uint8, dy)

	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			s[y] = append(s[y], uint8(picture2(x, y)))
		}
	}

	return s
}

func picture1(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func picture2(x, y int) int {
	return (x+y) / 2
}

func main() {
	pic.Show(Pic)
}
