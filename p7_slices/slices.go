// Slices are commong used because in different of arrays they have allow length changes
// What we mean with size changes, is that if initialy we have a slice of 10 values of an array of size 100 we can change later to 20 and will take the 10 next values
// Is like other languages is a peaces of a basis array
// Important the size of a slice can be <= to the size of the basis array

// Pointer decide the start and Length is how many values we will have
// Capacity is the maximun values allowed in the slice
package main
import "fmt"


func main() {
	arr := [...]string{"a", "b", "c", "d", "e", "f", "g"}

	s1 := arr[1:3] // end - 1
	s2 := arr[2:4]

	// show slices, length and capacity
	fmt.Println(s1)
	fmt.Println(s2)

	fmt.Println(len(s1), cap(s1))

	// Literal slice
	slice := []int{1, 2, 3, 4}
	fmt.Println(slice)

	// Variable slice
	// You can use make(data type, length, capacity)
	// also make(data type, length and capacity)
	slice_variable := make([]int, 8) // length = capacity = 8
	fmt.Println(slice_variable)

	slice_variable_2 := make([]int, 0, 3)
	slice_variable_2 = append(slice_variable_2, 4)
	fmt.Println(slice_variable_2)
}