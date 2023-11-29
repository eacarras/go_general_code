package main
import "fmt"


// Array is a fixed-length series of elements of choosen type this helps to the compiler to know 
// how much space they need to allocate
// In creation arrays always been initialized with the zero value of the the data type choosen
func main() {
	var x[5] int

	x[0] = 10
    fmt.Println(x[1]) // should print 0
	fmt.Println(x[0]) // should prit 10

	// array literal = array initialized
	y := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(y)

	// other way infering the length of the array
	z := [...]int{2, 3, 5, 7}
	fmt.Println(z)

	fmt.Println("Iterated over the array Z")
	for idx, value := range z {
		fmt.Println(idx, value)
	}
}
