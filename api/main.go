package main

import (
	"fmt"

	"github.com/ajagnic/apartment-site/server"
)

func main() {
	fmt.Println("Hello")

	test := server.Test()
	fmt.Println(test)
}
