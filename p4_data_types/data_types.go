package main

// IMPORTANT: This file doesnt work because there are multiple variables not used
func main() {
	// Maybe some time we need to convert data types for any reason
	// Let's see an example
	var x int32 = 1
	var y int16 = 2

	// Something wrong is try x = y
	// To convert we use T() where T is the data type
	x = int32(y)

	// Let's go with floats, remember that floats belongs to the R set
	// float32 = 6 digit precision , float64 = 15 digit precision
	var fx float32 = 132.45
	var fy float64 = 1.2345e2

	// We can not convert fx into fy because both have different space in memory

	const x = 3 // in constants the compiler can infer the data type with the value of the right
	// for enumerates you can use iota instead of just constants remeber when you use iota you not care about the value
	type Grades int
	const (
        A Grades = iota
        B
        C
        D
	)
}