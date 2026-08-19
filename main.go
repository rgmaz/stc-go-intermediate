package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 2)
	go func() {
		time.Sleep(30 * time.Millisecond)
		ch <- "slow"
	}()
	go func() {
		time.Sleep(10 * time.Millisecond)
		ch <- "fast"
	}()

	select {
	case v := <-ch:
		fmt.Println(v)
	case <-time.After(time.Millisecond * 100):
		fmt.Println("timeout")
	}
}
