package main

const s = "Go101.org"

// len(s) == 9
// 1 << 9 == 512
// 512 / 128 == 4

func main() {
	sizea := len(s)
	sizeb := len(s[:])
	var a byte = 1 << sizea / 128
	var b int = 1 << sizeb
	println(b)
	println(len(s[:]))
	println(a, b)
}
