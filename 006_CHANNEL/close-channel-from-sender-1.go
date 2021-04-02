package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan bool)
	timeout := time.After(600 * time.Millisecond)
	go sendAndClose(ch)

	for {
		select {
		case <-ch:
			fmt.Println("Got message")
		case <-timeout:
			fmt.Println("Timeout")
			return
		default:
			fmt.Println("*yawn*")
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func sendAndClose(ch chan bool) {
	time.Sleep(120 * time.Millisecond)
	ch <- true
	close(ch)
	fmt.Println("Send and closed")
}
