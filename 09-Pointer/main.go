package main

import "fmt"

type myStruct struct {
	foo int
}

func main() {
	var a int = 42
	var b *int = &a
	fmt.Println(a, *b)
	a = 50
	fmt.Println(a, *b)
	*b = 14
	fmt.Println(a, *b)

	d := [3]int{1, 2, 3}
	e := &d[0]
	f := &d[1]
	fmt.Printf("%v %p %p\n", d, e, f)

	var ms *myStruct
	fmt.Println(ms)
	ms = new(myStruct)
	fmt.Println(ms)
	(*ms).foo = 40
	fmt.Println(ms.foo)

	//array -> copy
	//slice, map -> point

}
