package main

import (
	"fmt"
	"math/rand"
)

func printRandom() {
	for i := 0; i < 2; i++ {
		fmt.Println(rand.Intn(100))
	}
}

func main() {

	for i := 0; i < 100; i++ {
		printRandom()
	}

}
