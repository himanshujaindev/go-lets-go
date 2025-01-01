package main // Entrypoint
import (
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"
)

type gasEngine struct {
	mpg     uint8
	gallons uint8
	owner
	int
}
type owner struct {
	name string
}

type electicEngine struct {
	mpkwh uint8
	kwh   uint8
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

func (e electicEngine) milesLeft() uint8 {
	return e.mpkwh * e.kwh
}

type engine interface {
	milesLeft() uint8 // method signature, object must have this defined
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it")
	} else {
		fmt.Println("Cant make it")
	}
}

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

	// switch
	switch {
	case num1 == 1:
		fmt.Println("num1 is 1")
	default:
		fmt.Println("default")
	}

	switch num1 {
	case 1:
		fmt.Println("num1 is 1")
	default:
		fmt.Println("default")
	}

	// array
	var intArr [3]int32 // [0,0,0] - 12 bytes
	intArr[1] = 1
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])

	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	// var intArr1 [3]int32 = [3]int32{2,3,4}
	// intArr1 := [3]int32{2,3,4}
	intArr1 := [...]int32{2, 3, 4, 4, 5, 6}
	fmt.Println(intArr1)

	// slice
	var intSlice []int32 = []int32{10, 20, 30}
	fmt.Printf("Length is %v with capacity %v\n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 40)
	fmt.Printf("Length is %v with capacity %v\n", len(intSlice), cap(intSlice)) // [10 20 30 40 * *]

	appendSlice := []int32{50, 60}
	intSlice = append(intSlice, appendSlice...)

	fmt.Println(intSlice)

	intSlice1 := make([]int32, 3, 8) // (type, len, cap)
	fmt.Println(intSlice1)

	// map[string]uint8 - key=string; value=uint8
	myMap := make(map[string]uint8)
	fmt.Println(myMap)

	myMap2 := map[string]uint8{"H": 1, "J": 2}
	fmt.Println(myMap2["H"])
	fmt.Println(myMap2["K"]) // value = default of type uint8 = 0

	var age, ok = myMap2["J"]
	fmt.Println(age, ok)

	delete(myMap2, "J")
	fmt.Println(myMap2)

	// Loops
	for name, age := range myMap2 { // order is not preserved
		fmt.Printf("Name: %v; Value: %v\n", name, age)
	}

	for i, v := range intArr1 {
		fmt.Printf("index: %v; value: %v\n", i, v)
	}

	// while loop in go
	i := 0
	for i < 5 {
		fmt.Println(i)
		i += 1
	}

	j := 0
	for {
		if j >= 5 {
			break
		}
		fmt.Println(j)
		j += 1
	}

	for k := 0; k < 5; k++ {
		fmt.Println(k)
	}

	// string
	var myString = "résumé"
	// utf-8 = [01110010, 11000011, 10101001, 01110011, 01110101, 01101101, 11000011, 10101001]
	// [r, é, * , s, u, m, é, *]
	// [114, 195, 169, 115, 117, 109, 195, 169]
	var indexed = myString[2] // 195
	fmt.Println(myString)
	fmt.Println(len(myString))
	fmt.Printf("%v, %T\n", indexed, indexed)

	for i, v := range myString { // intellegent = é = 233
		fmt.Println(i, v)
	}

	// easy way = cast into rune
	var myRune1 = []rune("résumé") // [114,233,115,117,109,233]
	fmt.Println(myRune1)
	fmt.Println(len(myRune1))

	strSlice := []string{"h", "i", "m"}
	// var catStr = ""
	var strBuilder strings.Builder
	for i := range strSlice {
		// catStr += strSlice[i] // creates new string everytime
		strBuilder.WriteString(strSlice[i])
	}
	// fmt.Printf("%v\n", catStr)
	fmt.Printf("%v\n", strBuilder.String())

	// struct
	// var myEngine gasEngine = gasEngine{mpg: 25, gallons: 15}
	var myEngine gasEngine = gasEngine{25, 15, owner{"Alex"}, 1}
	myEngine.mpg = 10
	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.name, myEngine.int)
	fmt.Println(myEngine.milesLeft())

	var myEngine2 = struct {
		mpg     uint8
		gallons uint8
	}{25, 10}

	fmt.Println(myEngine2.mpg, myEngine2.gallons)

	// interface
	canMakeIt(myEngine, 50)

	var myEngine3 electicEngine = electicEngine{25, 15}
	canMakeIt(myEngine3, 250)

	// pointers
	var p *int32 = new(int32) // depends on OS; 8 bytes; default=nil
	*p = 10                   // set the value of memory location p is pointing to, to 10
	var ii int32
	fmt.Printf("value of p = %v\n", *p) // de-referencing pointers
	fmt.Printf("addr store by p = %v\n", p)
	fmt.Printf("location of p = %v\n", &p)
	fmt.Printf("value of i = %v\n", ii)

	p = &ii // mem address of ii
	*p = 100
	fmt.Printf("value of p = %v\n", *p) // de-referencing pointers
	fmt.Printf("value of i = %v\n", ii)
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
