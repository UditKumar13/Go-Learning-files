package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
	"unicode/utf8"
)

func printMe(printValue string) {
	fmt.Println(printValue)
}

func calculate(a int, b int) (int, int) {
	var sum int
	sum = a + b
	multiply := a * b
	return sum, multiply
}

func divide(a int, b int) (float64, int, error) {
	var err error
	if b == 0 {
		err = errors.New("division by zero is not allowed")
		fmt.Println(err)
		return 0, 0, err // Return zero values if division by zero
	}
	return float64(a) / float64(b), a % b, err
}

func main() {
	fmt.Println("Hello World!")
	var intNum_1 int64 = 4
	fmt.Println("intNum_1:", intNum_1)
	var intNum_2 int64 = 3
	fmt.Println("intNum_2:", intNum_2)

	fmt.Println("it will give only floor value", intNum_1/intNum_2) // No need for type conversion

	var floatNum float64 = 3.14
	fmt.Println(floatNum)
	var str string = "Hello, Go!"
	fmt.Println(str)

	// String length in bytes
	fmt.Println("Length of str in bytes:", len(str))
	// String length in runes (characters)
	fmt.Println("Length of str in runes:", utf8.RuneCountInString(str))
	// String length in runes (characters) using range
	count := 0
	for range str {
		count++
	}
	fmt.Println("Length of str in runes (using range):", count)
	//concatenation
	str2 := " How are you?"
	fmt.Println("Concatenated string:", str+str2)

	var myBoolean bool = true
	fmt.Println("Boolean value:", myBoolean)
	var myRune rune = 'A'
	fmt.Println("Rune value:", myRune)
	fmt.Println("Rune as string:", string(myRune))

	//assigning in a cool way

	var1 := 10
	var2 := "Udit is a software developer"
	fmt.Println("var1:", var1, "var2:", var2)

	// creating const

	const myConst string = "This is a constant string"
	fmt.Println("Constant value:", myConst)
	// myConst = "This will not work" // Uncommenting this line will cause a compilation error because constants cannot be reassigned

	var printValue string = "This is a print value from a parameteric function"
	printMe(printValue) // Calling the function with a string argument

	var sum, multipart int = calculate(5, 10) // Calling the function with two integers
	fmt.Println("Sum of 5 and 10 is:", sum)
	fmt.Println("Multiplication of 5 and 10 is:", multipart)
	fmt.Printf("Sum of 5 and 10 is %v and Multiplication is %v\n", sum, multipart)

	// Division function
	// handling errors also
	divisionResult, remainder, err := divide(10, 2)
	if err != nil {
		fmt.Printf(err.Error())
	} else if remainder == 0 {
		fmt.Println("Division of 10 by 2 is:", divisionResult)
	} else {
		fmt.Println("Division of 10 by 2 is:", divisionResult, "with a remainder of:", remainder)
	}

	// arrays sclices and maps usuage

	// arrays : fixed length, same type elements, contiguous memory allocation
	var intArray [5]int32 = [5]int32{1, 2, 3, 4, 5}
	fmt.Println("Integer Array:", intArray)

	fmt.Println(intArray[1:3]) // Slicing the array to get elements from index 1 to 2
	intSlice := []int32{6, 7, 8}
	fmt.Println("Slice of Integer Array:", intSlice)
	fmt.Printf("Length of array is %v and cap is %v\n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 9) // Appending an element to the slice
	fmt.Println("Slice after appending:", intSlice)
	fmt.Printf("Length of array is %v and cap is %v\n", len(intSlice), cap(intSlice))

	// we will make a slice using make function
	intSlice2 := make([]int32, 3, 5) // Create a slice with length 3 and capacity 5
	fmt.Println("Slice created using make:", intSlice2)
	fmt.Println("Length of intSlice2 is", len(intSlice2), "and capacity is", cap(intSlice2))

	// Maps: key-value pairs, unordered, dynamic size
	intMap := make(map[string]int)
	intMap["one"] = 1
	intMap["two"] = 2
	intMap["three"] = 3
	fmt.Println("Integer Map:", intMap)

	for key, value := range intMap {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}

	// order is not guaranteed in maps, so the output order may vary

	// check value exist in a map or not
	value, exists := intMap["two"]
	if exists {
		fmt.Printf("Value for key 'two': %d\n", value)
	} else {
		fmt.Println("Key 'two' does not exist in the map")
	}

	//delete a key from map
	delete(intMap, "two") // it deletes the key "two" from the map by reference
	fmt.Println("Map after deleting key 'two':", intMap)

	for i := 0; i < 5; i++ {
		fmt.Println("This is a loop iteration:", i)
	}
	// same thing can be done using for range loop
	for index, value := range intSlice {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// Using time package to get current time
	currentTime := time.Now()
	fmt.Println("Current Time:", currentTime)

	var n int = 1000000
	var testSlice = []int{}            // without pre-allocating memory
	var testSlice2 = make([]int, 0, n) // Pre-allocate memory for the slice

	// let us see the time taken to append elements to the slice
	startTime := time.Now()
	for i := 0; i < n; i++ {
		testSlice = append(testSlice, i)
	}
	elapsedTime := time.Since(startTime)
	fmt.Println("Time taken to append elements (without pre-allocation):", elapsedTime)

	// Now, let's do the same with the pre-allocated slice
	startTime = time.Now()
	for i := 0; i < n; i++ {
		testSlice2 = append(testSlice2, i)
	}
	elapsedTime = time.Since(startTime)
	fmt.Println("Time taken to append elements (with pre-allocation):", elapsedTime)
	// The pre-allocated slice should be faster as it reduces the need for multiple memory allocations

	// string manipulation
	str3 := "Hello, World!"
	fmt.Println("Original String:", str3)
	str3 = str3 + " How are you?"
	fmt.Println("Modified String:", str3)

	// String to rune conversion
	runeSlice := []rune(str3)
	fmt.Println("String to Rune Slice:", runeSlice)
	fmt.Println("Rune Slice as String:", string(runeSlice))

	// more on string :

	var mystring = "a resume"
	var mystring2 = []rune("a resume")
	var indexed = mystring[0] // This will give the byte value of the first character
	fmt.Println("First character byte value:", indexed)

	fmt.Printf("%v, %T\n", indexed, indexed) // %v for value, %T for type

	fmt.Printf("value : %v, type : %T\n", mystring, mystring) // %v for value, %T for type
	for i, v := range mystring {
		fmt.Println(i, v) // Using fmt.Sprintf to get type as string
	}

	fmt.Printf("value : %v, type : %T\n", mystring2, mystring2) // %v for value, %T for type

	for i, v := range mystring2 {
		fmt.Println(i, v) // Using fmt.Sprintf to get type as string
	}

	// strings are immutable in Go, so you cannot change a character in a string directly
	// mystring[0] = 'b' // This will cause a compilation error

	fmt.Println("we can use string builder to modify a string")
	var strBuilder strings.Builder
	strBuilder.WriteString("Hello")
	strBuilder.WriteString(", ")
	strBuilder.WriteString("World!")
	fmt.Println("Modified String:", strBuilder.String())

	fmt.Println("To convert back to string from string builder we can use String() method of strings.Builder")
	var converted_string = strBuilder.String()
	fmt.Println("Converted String:", converted_string)

}
