package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	fields := strings.Fields(sc.Text())
	a := make(chan int)
	b := make(chan int)
	// Producer
	go func() {
		defer close(a)
		for _, f := range fields {
			n, _ := strconv.Atoi(f)
			a <- n
		}
	}()
	// Squarer
	go func() {
		defer close(b)
		for n := range a {
			b <- (n * n)
		}
	}()
	sum := 0
	for v := range b {
		sum += v
	}
	fmt.Println(sum)
}
