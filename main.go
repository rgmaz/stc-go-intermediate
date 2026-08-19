package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		nums[i], _ = strconv.Atoi(sc.Text())
	}

	var mu sync.Mutex
	var wg sync.WaitGroup
	total := 0

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
	// TODO: split `nums` into 4 chunks and launch a goroutine per chunk.
	// Each goroutine should sum its chunk and add the result into `total`,
	// guarded by `mu`. Use `wg` (Add/Done/Wait) to wait for all goroutines
	// to finish before printing.

	wg.Wait()
	fmt.Println(total)
}
