package main
import "fmt"

// One line comment
/*
 * This is a general comment
*/
func main() {
	var name string = "Aaron"
	fmt.Println("Hi: " + name)
	fmt.Println("Hi:", name) // Same way add a space before the value

	var number int = 2
	fmt.Println("Number:", number)
	var sum int = 2 * 2 + 1
	fmt.Println("Total:", sum)
}
