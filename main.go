package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func basic() {
	fmt.Println("Hello World!")
}

func variableDeclaration() {
	name := "John"
	var s string = "Hello World!"
	var i int = 42
	var b bool = true
	var f float64 = 3.14

	fmt.Println(s)
	fmt.Println(i)
	fmt.Println(b)
	fmt.Println(f)
	fmt.Println(name)

}

func printVariableString() {

	var name string = "KodeKloud"
	var user string = "John"
	fmt.Print("Hello ", name, ",", user)
}

func printFormatSpecifier() {

	var name string = "KodeKloud"
	var i int = 42

	// fmt.Printf("Hello %v", name)
	fmt.Printf("Hello %v, you are %d years old", name, i)
}

func variableDeclarationAndInitialization() {
	var name string
	var i int
	var b bool
	var f float64

	name = "John"
	i = 42
	b = true
	f = 3.14

	fmt.Println(name)
	fmt.Println(i)
	fmt.Println(b)
	fmt.Println(f)

	// Short hand way of declaring and initializing a variable
	var s, t string = "Hello", "World"
	fmt.Println(s, t)

	var (
		status string = "OK"
		age    int    = 42
	)
	fmt.Println(status, age)

	// Short variable declaration
	short := "Short variable declaration"
	fmt.Println(short)
}

func scope() {

	name := "John"
	{
		var name string = "Doe"
		fmt.Println(name)
	}
	fmt.Println(name)
}

func zeroValue() {
	var name string
	var i int
	var b bool
	var f float64

	fmt.Println(name)
	fmt.Println(i)
	fmt.Println(b)
	fmt.Println(f)
}

func singleInput() {
	var name string
	fmt.Printf("Type your name: ")
	fmt.Scanf("%s", &name)
	fmt.Println("Hey there, ", name)
}

func multipleInput() {
	var name string
	var age int
	fmt.Printf("Type your name and age: ")
	count, err := fmt.Scanf("%s %d", &name, &age)
	fmt.Println("Count: ", count)
	fmt.Println("Error: ", err)
	fmt.Printf("Hey there, %s, you are %d years old", name, age)
}

func typeOfVariable() {
	var name string = "John"
	var i int = 42
	var b bool = true
	var f float64 = 3.14

	fmt.Printf("Type of name: %T\n", name)
	fmt.Printf("Type of i: %T\n", i)
	fmt.Printf("Type of b: %T\n", b)
	fmt.Printf("Type of f: %T\n", f)

	data := reflect.TypeOf(name)
	fmt.Println(data.Name())
}

func typeCasting() {

	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)

	fmt.Printf("Type of i: %T and value is %v\n", i, i)
	fmt.Printf("Type of f: %T and value is %.2f\n", f, f)
	fmt.Printf("Type of u: %T and value is %v\n", u, u)

	// Integer to String
	var s string = strconv.Itoa(i)
	fmt.Printf("%q", s)

	// String to Integer
	var str string = "42"
	j, err := strconv.Atoi(str)

	fmt.Printf("Type of j: %T and value is %v\n", j, j)
	fmt.Printf("Type of err: %T and value is %v\n", err, err)
}

func main() {

	typeCasting()
}
