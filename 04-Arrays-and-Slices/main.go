package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("----- Basic arrays -----")
	grades := [3]int{98, 89, 85}
	fmt.Println(grades)

	var students [5]string
	fmt.Println(students)
	students[0] = "Mixko"
	students[1] = "Apisit"
	students[2] = "GG"
	fmt.Printf("Students: %v\n", students)

	fmt.Printf("Number of students: %v\n", len(students))
	fmt.Println("----- -----")

	a := []int{1, 2, 3}
	b := a // re-assign
	// b := &a // point to the same data
	b[1] = 8
	fmt.Printf("%v\n%v\n", a, b)

	fmt.Println("----- Slice -----")

	c := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	d := c[:]   //Slice of all element
	e := c[3:]  //Slice from 4th to end
	f := c[:6]  //Slice from the first 6 elements
	g := c[3:6] //Slice from 4-6th elements
	c[5] = 42
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(len(c), cap(c))

	fmt.Println("----- Slice#2 -----")
	h := make([]int, 3, 100)
	fmt.Println(h)
	fmt.Printf("Length: %v\n", len(h))   //tells you how many items are in the array.
	fmt.Printf("Capacity: %v\n", cap(h)) // tells you the capacity of the underlying array.
	fmt.Println("----- Slice#3 -----")
	i := []int{}
	i = append(i, 1)
	i = append(i, 2, 3, 4, 5, 6, 7)

	fmt.Println(i)
	fmt.Printf("Length: %v\n", len(i))
	fmt.Printf("Capacity: %v\n", cap(i))
	j := append(i[:3], i[6:]...)
	fmt.Println(j)
	fmt.Printf("Length: %v\n", len(j))
	fmt.Printf("Capacity: %v\n", cap(j))

	k := []int{12, 234, 756, 2, 5, 73, 4, 6, 456, 675}
	sort.Ints(k)
	fmt.Println(k)
}
