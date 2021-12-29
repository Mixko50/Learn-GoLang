package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("start")
	// defer fmt.Println("middle")
	// fmt.Println("end")
	// panic("Hello")

	defer func() {
		if err := recover(); err != nil {
			log.Println("Error: ", err)
		}
	}()
	panic("Something bad happend")
}
