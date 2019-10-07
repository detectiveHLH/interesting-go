package main

import "fmt"

type Coord struct {
	X int32
	Z int32
}

func main() {
	fmt.Println(test())
}

func test() []Coord {
	var test []Coord
	for i := 0; i < 5; i++ {
		test = append(test, Coord{X:2,Z:2})
	}
	return test
}