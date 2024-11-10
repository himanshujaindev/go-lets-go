package main // Entrypoint
import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("Hello")

	// int8 / int16 / int32 / int64 / uint8 / uint16 / uint32 / uint64
	// uint -> only pos num
	var intNum int = 10
	fmt.Println(intNum)

	var floatNum float64 = 3.142425
	fmt.Println(floatNum)

	// Cant add / mul - float and int
	// To do this - Casting
	// Div / -> Round the res

	var myString1 string = "Hello"
	var myString2 string = `Hello
	World`
	fmt.Println(myString1 + myString2)
	fmt.Println(len(myString1))                    // Num of bytes
	fmt.Println(utf8.RuneCountInString(myString1)) // Len of string

	var myRune rune = 'a' // Unicode codepoint
	fmt.Println((myRune))

	var myBool bool = true
	fmt.Println(myBool)

	// default: all int uint float = 0 , string = '' , bool = false
	var intNum1 int // default = 0
	fmt.Println(intNum1)

	myVar := "Hello" // var myVar = "Hello" // Type is inferred
	fmt.Println(myVar)

	num1, num2 := 1, 2
	fmt.Println(num1 + num2)

	const myConst string = "const" // Cant change after created ; Initial
	fmt.Println(myConst)

	var myReturn1, myReturn2, err = printMe("Hi")
	if err != nil { // && , ||
		fmt.Println(err.Error())
	} else if myReturn2 == 0 {
		fmt.Printf("%v - %v", myReturn1, myReturn2)
	}

}

func printMe(printVal string) (int, int, error) {
	fmt.Println(printVal)

	var err error // nil
	num1, num2 := 1, 0
	if num2 == 0 {
		err = errors.New("div by zero error")
		return 0, 0, err
	}
	var divNum = num1 / num2
	fmt.Println(divNum)
	return 1, 1, err
}
