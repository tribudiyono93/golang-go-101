package main

var x, y, z = 123, true, "foo"

func main() {
	var q, r = 789, false
	r, s := true, "bar"
	r = y
	x = q

	_, _ = r, s
}
