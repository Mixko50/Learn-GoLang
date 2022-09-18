package main

import (
	"fmt"
	"re-skill/customer"
)

func main() {
	x := customer.Person{}
	x.SetName("Mixko")
	fmt.Printf("%#v", x.GetName())
}
