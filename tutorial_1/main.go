package main

import (
	"errors"
	"fmt"
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
}

// The divide function handles division and returns an error if division by zero is attempted
