package main

import "fmt"

var msg chan interface{}

func main() {
	msg = make(chan interface{}, 5)

	for i := 0; i < 2; i++ {
		send(i)
	}


	for {
		select {
		case info := <- msg:
			fmt.Printf("receive %v", info)
			fmt.Println()
		}
	}

	for i := 0; i < 5; i++ {
		send(i)
	}


}

func send(info interface{}) {
	select {
	case msg <- info:
	default:
		fmt.Println("full")
	}
}
