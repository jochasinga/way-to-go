package main

var a = "G"

func main() {
	n() // print G
	m() // print O
	n() // print G
}

func n() { print(a) }
func m() {
	a := "O"
	print(a)
}
