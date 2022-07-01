package main

import (
	"fmt"
)

func main() {
	g := []string{"Me", "You", "Them"}

	for i, b := range g {
		fmt.Println(i, b)
	}
}
