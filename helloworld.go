package gomodhelloworld

import "fmt"

func HelloWorld() string {
	return "Hello Dunia"
}

func SayHello(name string) string {
	fmt.Println("Hello", name)
	return name
}
