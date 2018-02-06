package main

import (
	"fmt"

	"github.com/KiraFox/gogal-dynamic/rand"
)

func main() {
	// Testing our String function given 10 bytes
	fmt.Println(rand.String(10))
	// Testing our RememberToken function which uses a constant
	fmt.Println(rand.RememberToken())
}
