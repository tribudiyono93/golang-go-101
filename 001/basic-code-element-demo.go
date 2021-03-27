package main

import "math/rand"

const MaxRnd = 16

func StatRandomNumbers(n int) (int, int) {
	var a, b int
	for i := 0; i < n; i++ {
		if rand.Intn(MaxRnd) < MaxRnd / 2 {
			a = a + 1
		} else {
			b++
		}
	}

	return a, b
}
func main() {
	var num = 100

	x, y := StatRandomNumbers(num)
	print("Result: ", x, " + ", y, " = ", num, "? ")
	println(x + y == num)
}
