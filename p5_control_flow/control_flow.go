package main
import "fmt"


func main() {
	// Control structure if
	var x int = 5
	if (x > 4) {
		fmt.Println("X is higher than 4")
	} else {
		fmt.Println("X is lower than 4")
	}
	// We can use also else if statements like other languages

	// Control structure for
	for i := 0; i < x; i++ {
		fmt.Println(i)
	}
	/**
	Other ways of for
	x := 1
	for x<5 {
		fmt.Println(x)
		x++
	}

	A fun way to create a infinite loop
	for {
        fmt.Println("Hello")
    }
	**/

	// Control structure switch
	switch x {
		case 3:
			fmt.Println("X is Three")
			// We dont need a break here, automatically break at the end of each case
		case 4:
			fmt.Println("X isFour")
		case 5:
			fmt.Println("X is Five")
		default:
			fmt.Println("Not a number")
    }

	// Control structure switch tagless
	switch {
		case x < 4:
			fmt.Println("X is lower than 4")
        case x > 4:
			fmt.Println("X is higher than 4")
		default:
			fmt.Println("Not a number")
	}

	// Remeber break and continue still exists here use when you think is necessary
}
