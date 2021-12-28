package main

import "fmt"

func main() {
	fmt.Println("--- Default value ---")
	n := 1 == 1
	m := 1 == 2
	fmt.Printf("%v, %T\n", n, n)
	fmt.Printf("%v, %T\n", m, m)

	//Default of boolean is false
	var boo bool
	fmt.Printf("%v, %T\n", boo, boo)

	//Default of string is empty
	var str string
	fmt.Printf("%v, %T\n", str, str)

	//Default of int and float are 0
	var num int
	fmt.Printf("%v, %T\n", num, num)
	var flo float32
	fmt.Printf("%v, %T\n", flo, flo)

	fmt.Println("--- Operation ---")
	var num1 int = 10
	var num2 int = 5
	fmt.Println(num1 + num2)
	fmt.Println(num1 - num2)
	fmt.Println(num1 + num2)
	fmt.Println(num1 / num2)
	fmt.Println(num1 % num2)

	fmt.Println("--- Convert type ---")
	var num3 int = 3
	var num4 int16 = 20 //different type => error
	fmt.Println(num4 / int16(num3))

	fmt.Println("--- Bit Operation ---")
	//Bit operation
	bit1 := 10                //1010
	bit2 := 3                 //0011
	fmt.Println(bit1 & bit2)  //0010
	fmt.Println(bit1 | bit2)  //1011
	fmt.Println(bit1 ^ bit2)  //ExOR 1001
	fmt.Println(bit1 &^ bit2) //Opposite of or 0100

	fmt.Println("--- String ---")
	s := "Hello world"
	fmt.Printf("%v, %T\n", s[2], s[2])
	fmt.Printf("%v, %T\n", string(s[2]), s[2])
	b := []byte(s)
	t := "test"
	fmt.Printf("%v, %T\n", b, b)
	fmt.Println(s, t)

	fmt.Printf("%q\n", t) //Add quote
}
