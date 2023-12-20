package main
import "fmt"

// HashTable: key-value pair => May have collisions
// Not ints in the key and not slices or arrays, is faster O(n)
func main() {
	// Initialize the hash table
	var idMap map[string]int
	idMap = make(map[string]int)

	fmt.Println(idMap) // empty map

	// Assign values to the map
	idMap["Pepe"] = 200
	fmt.Println(idMap) // { "Pepe": 200 }

	// Delete key/value
	idMap["Julio"] = 100
	fmt.Println(idMap) // { "Pepe": 200, "Julio": 100 }
	delete(idMap, "Julio")
	fmt.Println(idMap) // { "Pepe": 200 }

	// Get id/value
	val, exists := idMap["Pepe"]
	val2, exists2 := idMap["Julio"]
	fmt.Println(val, exists) // 200, true
	fmt.Println(val2, exists2) // 0, false (the first one is the default value of the type)

	// Iterate over the map
	idMap["Julio"] = 100
	idMap["Maria"] = 195

	for key, val := range idMap {
		fmt.Println(key, val)
	}
}
