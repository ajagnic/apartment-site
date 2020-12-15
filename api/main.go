package main

import (
	"fmt"

	"github.com/ajagnic/apartment-site/server"
)

func main() {
	err := server.Run(":8080")
	fmt.Println(err)
}
