// Write a program which prompts the user to enter a string. The program searches through the entered string for the characters ‘i’, ‘a’, and ‘n’.
// The program should print “Found!” if the entered string starts with the character ‘i’, ends with the character ‘n’, and contains the character ‘a’.
// The program should print “Not Found!” otherwise. The program should not be case-sensitive, so it does not matter if the characters are upper-case or lower-case.
package main
import (
	"fmt"
    "strings"
)


func main() {
	var s string
	
	fmt.Print("Enter a string to validate: ")
	_, err :=  fmt.Scan(&s)

	if err != nil {
		fmt.Println(err)
	}

	// First make all lowe case
	s = strings.ToLower(s)
	
	// Validations
	startsWith := strings.HasPrefix(s, "i")
	endsWith := strings.HasSuffix(s, "n")
	containsA := strings.Contains(s, "a")

	if startsWith && endsWith && containsA {
		fmt.Println("Found")
	} else {
		fmt.Println("Not Found")
	}
}