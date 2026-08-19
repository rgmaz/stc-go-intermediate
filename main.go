package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)

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

	subLen := len(nums) / 4
	tasks := [][]int{
		nums[:subLen],
		nums[subLen : subLen*2],
		nums[subLen*2 : subLen*3],
		nums[subLen*3:],
	}

	for _, task := range tasks {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sum := 0
			for _, num := range task {
				sum += num
			}
			mu.Lock()
			total += sum
			mu.Unlock()
		}()
	}
	// TODO: split `nums` into 4 chunks and launch a goroutine per chunk.
	// Each goroutine should sum its chunk and add the result into `total`,
	// guarded by `mu`. Use `wg` (Add/Done/Wait) to wait for all goroutines
	// to finish before printing.

	wg.Wait()
	fmt.Println(total)
}
