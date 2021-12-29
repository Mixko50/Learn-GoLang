package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	i := 1
	for ; i < 5; i++ {
		fmt.Println(i)
	}

	s := []int{1, 2, 3}
	for k, v := range s {
		fmt.Println(k, v)
	}

	str := "Hello Go!"
	for k, v := range str {
		fmt.Println(k, string(v))
		break
	}

	statePopu := map[string]int{
		"Cali": 345,
		"Tex":  100,
		"TO":   23,
		"New":  1234,
		"Ohio": 213,
	}

	for _, v := range statePopu {
		fmt.Println(v)
	}
}
