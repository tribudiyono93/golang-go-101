package main

func main() {
	const (
		k = 3
		m float32 = iota + .5
		n
		p = 9
		q = iota * 2
		_
		r
		s, t = iota, iota
		u, v
		_, w
	)
	const x = iota
	const (
		y = iota
		z
	)

	print(m)
	println(n)
	println(q, r)
	println(s, t, u, v, w)
	println(x, y, z)
}
