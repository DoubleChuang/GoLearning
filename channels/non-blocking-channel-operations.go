package main

import (
	"fmt"
)

func main() {
	messages := make(chan string, 1)
	signals := make(chan bool)
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}
	msg := "Hi"

	select {
	case messages <- msg:
		fmt.Println("Sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}

}
