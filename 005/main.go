package main

const y = 789
var x int = 123

func main() {
	var x = true

	{
		x, y := x, y
		x, z := !x, y/10
		y /= 100

		println(x, y, z)
	}
	println(x)
	//println(z)
}
