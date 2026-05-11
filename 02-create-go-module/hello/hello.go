package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Hector")
	fmt.Println(message)
}
