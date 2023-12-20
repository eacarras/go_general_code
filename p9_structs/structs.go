package main
import "fmt"

type Person struct {
	name    string
	address string
	phone   string
}
// Struct is a custom object that represents something in the real world
func main() {
	var person1 Person

	// Use dot notation to add values or get
	person1.name = "Pepe"
	var x string = person1.name

	fmt.Println(x)

	// Initialize with field to zero
	var person2 *Person
	person2 = new(Person)

	fmt.Println(*person2)

	// Initialize with field not zero
	person3 := Person{name: "John", address: "Ultimate", phone: "1234" }
	fmt.Println(person3)
}
