package main

import "fmt"

func main() {
	var w Writer = consoleWriter{}
	w.Write([]byte("Hello, Mixko!"))
}

type Writer interface {
	Write(data []byte) (int, error)
}

type consoleWriter struct{}

func (cw consoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}
