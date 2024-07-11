package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	// Type: Integers

	var intNum int = 123 // we also have int8 int16 int32 int64 as the dtypes and by default int uses 32 or 64 bit depending on the system arch
	fmt.Println(intNum)

	var intNum16 int16 = 32767
	intNum16 = intNum16 + 1 // doesn't throw the run time error but this can produce wierd results
	fmt.Println(intNum16)

	// uint - unsigned int's doesn't store the negative values
	// because of this we can store integers twice as large in the same amount of memory
	// i.e,
	// Ex:
	// int8: (-128, 127)
	// uint8: (0, 255)

	// Type: Floats - used to store decimal numbers
	// essentially most floating point numbers won't be stored precisely on your computer like float32 declaration

	var floatNum float64 = 1313434.9 // go doesn't have just float type we need to specify either float32 or float64 as dtype
	fmt.Println(floatNum)

	// Note: 2 Things to note in airthmetic operations:

	//   1. we can't perform operations on the mixed types; sol: we need to cast one of the numbers ot the coommon type
	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)

	//   2. integer division results in an integer and the result is rounded down
	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1 / intNum2)
	fmt.Println(intNum1 % intNum2)

	// Type: String
	// we can initialize the value of a variable using double quotes "" (for single line) or back quotes `` (for multi-line)

	var myString string = "Hello \nWorld"
	var myString2 string = `Hello  
World`
	fmt.Println(myString)
	fmt.Println(myString2)

	// we can concatinate string by adding
	var myString3 = "Hello" + " " + "World"
	fmt.Println(myString3)

	// Imp: Go uses utf-8 encoding
	// len() gives the length of the number of bytes used for the string
	// As Go uses utf-8 encoding. The characters outside the vanilla ASCII character set (256) are stored with more than a single byte
	fmt.Println(len("A")) // >> 1
	fmt.Println(len("γ")) // >> 2

	// Hence to find the lenght of the string that may contain the special characters we can use the built-in package unicode/utf8 and use the function RuneCountInString()
	fmt.Println(utf8.RuneCountInString("γ")) // >> 1

	// runes's are another data types in Go and represent characters
	var myRune rune = 'a' // single quote character '' is called a rune
	fmt.Println(myRune)

	// Type: Booleans

	var myBoolean bool = true // or false
	fmt.Println(myBoolean)

	// Default Values

	// the default value for all ints uints floats and runes is 0
	// i.e, for all the below data types the default value if not initialized is 0
	// uint, uint8, uint16, uint32, uint864
	// int, int8, int16, int32, int64
	// float32, float64
	// rune
	var intNum3 int
	fmt.Println(intNum3) // defaults to 0

	// for string default value is ''
	// for bool the default value is false

	// storthand := and auto dtype infering
	myVar := "text"
	fmt.Println(myVar) // >> text

	// initializing multiple variables at once
	// var var1, var2 int = 1, 2
	var1, var2 := 1, 2
	fmt.Println(var1, var2)

	// adding the type when the dtype is not obvious is a good practice
	// for ex: below we do not know the return type of the foo(), hence declaring the dtype is advisable
	// myVar := foo()  // avoid this
	// var myVar string := foo()  // good practice

	// Constants
	// constants are alternatives to variables
	// everything we said about variables also applies to constants except we can't change its value once its created
	// also we can't just declare constants. we always have to initialize

	const myConst string = "const value"
	fmt.Println(myConst)

	// constants can be used for some constants values in the code like pi as shown below
	const pi float32 = 3.145
	fmt.Println(pi)

}
