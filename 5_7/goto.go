package main

import "fmt"

func main() {
	fmt.Println(0)
	goto A

E:
	fmt.Println(6)
	goto F
F:
	fmt.Println(7)
	goto G
G:
	fmt.Println(8)
	goto H
H:
	fmt.Println(9)
	goto I
I:
	fmt.Println(10)
	goto J
A:
	fmt.Println(1)
	goto B
B:
	fmt.Println(2)
	goto C
K:
	fmt.Println(12)
	goto L
L:
	fmt.Println(13)
	goto M
M:
	fmt.Println(14)
	return
C:
	fmt.Println(3)
	goto D
D:
	fmt.Println(4)
	goto E
J:
	fmt.Println(11)
	goto K
}
