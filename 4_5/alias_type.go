// Define an alias type Rope for string and declare a variable with it
package main
import "fmt"

type Rope string

var myVar Rope

func main() {
	myVar = "I used to be a string"
	fmt.Println(myVar)
}