package main

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			print("SKIPPED")
			continue
		}
		print(i)
		print(" ")
	}
}
