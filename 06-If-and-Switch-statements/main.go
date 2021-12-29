package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("----- if else -----")
	statePopu := map[string]int{
		"Cali": 345,
		"Tex":  100,
		"TO":   23,
		"New":  1234,
		"Ohio": 213,
	}

	if pop, ok := statePopu["Tex"]; ok {
		fmt.Println(pop)
	}

	if statePopu["Cali"] < 50 {
		fmt.Println("Wow")
	} else if statePopu["Cali"] > 300 {
		fmt.Println("Too high")
	} else {
		fmt.Println("sdsd")
	}
	fmt.Println(statePopu["Cali"] > 50) //true

	myNum := 0.123
	if math.Abs(myNum/math.Pow(math.Sqrt(myNum), 2)-1) < 0.01 {
		fmt.Println("same")
	} else {
		fmt.Println("diff")
	}
	fmt.Println("----- Switch -----")
	switch i := 2 + 3; i {
	case 1, 5:
		fmt.Println("It is 1 or 5")
	case 2:
		fmt.Println("It is 2")
	default:
		fmt.Println("Not 1 or 2")
	}

	j := 10
	switch {
	case j <= 10:
		fmt.Println("less than or equal 10")
		fallthrough //check the next case
	case j <= 20:
		fmt.Println("less than or equal 20")
	default:
		fmt.Println("Not 1 or 2 or 5")
	}

	var i interface{} = 2
	switch i.(type) {
	case int:
		fmt.Println("int")
	default:
		fmt.Println("Not int")
	}
}
