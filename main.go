package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func add(numbers []int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}

	return total
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	sc.Scan()

	fields := strings.Fields(sc.Text())
	nums := make([]int, n)
	for i, f := range fields {
		nums[i], _ = strconv.Atoi(f)
	}
	var wg sync.WaitGroup
	var total int
	var mu sync.Mutex

	subLen := n / 4
	chunks := [][]int{
		nums[:subLen],
		nums[subLen : subLen*2],
		nums[subLen*2 : subLen*3],
		nums[subLen*3:],
	}
	for _, i := range chunks {
		wg.Add(1)
		go func(n []int) {
			defer wg.Done()

			subTotal := add(n)
			mu.Lock()
			total += subTotal
			mu.Unlock()
		}(i)
	}
	wg.Wait()
	fmt.Println(total) // replace with the real total
}
