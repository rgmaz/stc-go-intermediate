package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const pi = 3.14

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return pi * (c.Radius * c.Radius)
}

type Square struct {
	Side float64
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

// type Shape interface { ... }
// type Circle struct { ... }
// func (c Circle) Area() float64 { ... }

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	kind := sc.Text()
	sc.Scan()
	dim, _ := strconv.ParseFloat(sc.Text(), 64)
	var s interface{ Area() float64 }
	switch kind {
	case "circle":
		s = Circle{Radius: dim}
	case "square":
		s = Square{Side: dim}
	}
	if s != nil {
		fmt.Printf("%.2f\n", s.Area())
	}
}
