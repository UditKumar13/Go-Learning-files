package main

import (
	"fmt"
	"unicode/utf8"
)

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

}
