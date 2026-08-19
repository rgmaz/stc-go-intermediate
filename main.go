package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Logger struct{}

func (l Logger) Log(msg string) {
	fmt.Printf("[log] %s\n", msg)
}

type Counter struct {
	Logger
	count int
}

func (c *Counter) Inc() {
	c.count++

	msg := strconv.Itoa(c.count)
	c.Log(msg)
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	var c Counter
	for range n {
		c.Inc()
	}
}
