// The idea here is that the user insert a float value and we need to print a that trunc digit
package main
import "fmt"


func main(){
	var f float64

	fmt.Print("Enter a float value: ")
	_, err := fmt.Scanf("%f", &f)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("You have entered: %d", int(f))
}