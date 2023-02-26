package main

var a, b int

func f() {
	a = 1 // w 之前的操作
	b = 2 // 写操作w
}

func g() {
	print(b) // 读操作r
	print(a) // ???
	println()
}

func main() {
	go f() // g1
	g()    // g2
}
