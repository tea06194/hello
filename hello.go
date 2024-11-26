package main

import (
	"fmt"

	"github.com/tea06194/greetings"
)

func main() {
	message, error := greetings.Hello("qwe")

	if error != nil {
		fmt.Println(error)
	}

	fmt.Println(message)
}