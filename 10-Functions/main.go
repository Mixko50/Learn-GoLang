package main

import "fmt"

func main() {
	greeting := "hello"
	name := "Mixko"
	s := sum(1, 2, 3, 4, 5)
	sayHi(&greeting, &name)
	fmt.Println(greeting, name)
	fmt.Println("The sum is", s)
	d, err := divide(5.0, 2.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)

	//Anonymous function
	var divide func(float64, float64) (float64, error)
	divide = func(f1, f2 float64) (float64, error) {
		if f2 == 0.0 {
			return 0.0, fmt.Errorf("Cannotr divide by zero")
		} else {
			return f1 / f2, nil
		}
	}

	e, err := divide(5.0, 1.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(e)

	//Struct
	f := greeter{
		greeting: "Hello",
		name:     "Mixko",
	}
	f.greet()
	fmt.Println(f.name)

}

type greeter struct {
	greeting string
	name     string
}

func (g *greeter) greet() {
	fmt.Println(g.greeting, g.name)
	g.name = "Apisit"
}

func sayHi(greeting, name *string) {
	fmt.Println(*greeting, *name)
	*name = "Apisit"
	fmt.Println(*greeting, *name)
}

func sum(values ...int) (result int) {
	for _, v := range values {
		result += v
		fmt.Println(v)
	}
	return
}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide bt zero")
	}
	return a / b, nil
}
