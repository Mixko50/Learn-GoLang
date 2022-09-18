package main

func main2() {
	c, _ := sum(1, 2)
	println(c)

	x := func(a, b int) int {
		return a + b
	}

	x2 := x(10, 20)
	println(x2)

	cal(add)
	cal(sub)
	cal(func(a, b int) int {
		return a + b
	})

	sum := sumArr(1, 2, 3, 4, 5, 6, 77, 8, 9, 0, 0)
	println(sum)
}

// Tuple is return multiple values
func sum(a, b int) (int, string) {
	return a + b, "Hello"
}

func cal(f func(int, int) int) {
	result := f(10, 20)
	println(result)
}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func sumArr(a ...int) int {
	total := 0
	for _, v := range a {
		total += v
	}
	return total
}
