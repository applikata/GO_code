package main

import (
	"fmt"
	"runtime"
	"time"
)

func printStats(mem runtime.MemStats) {
	runtime.ReadMemStats(&mem)
	fmt.Println("mem.Alloc: ", mem.Alloc)
	fmt.Println("mem.TotalAlloc: ", mem.TotalAlloc)
	fmt.Println("mem.HeapAlloc: ", mem.HeapAlloc)
	fmt.Println("mem.NumGC: ", mem.NumGC)
}

func main() {
	var mem runtime.MemStats
	printStats(mem)

	for i := 0; i < 10; i++ {
		s := make([]byte, 50000000)
		if s == nil {
			fmt.Println("Operator Failed!")
		}
	}

	printStats(mem)

	for i := 0; i < 10; i++ {
		s := make([]byte, 10000000)
		if s == nil {
			fmt.Println("Operator Failed!")
		}
		time.Sleep(5 * time.Second)
	}
	printStats(mem)
}
