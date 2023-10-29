package main
import "fmt"

func main() {
	// Create a variable and initialize with a integer variable
	var integer int = 2
	var ip *int = &integer // a pointer to integer

	var copy int = *ip // We copy the value from the pointer, getting the value allocated in that space
	fmt.Println(copy) // should be 2
}