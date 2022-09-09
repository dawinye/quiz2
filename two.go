package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

func main() {
	var intArg, _ = strconv.Atoi(os.Args[1])
	var sliceX []int
	for i := 0; i < intArg; i++ {
		sliceX = append(sliceX, rand.Intn(intArg))
	}
	var newSlice []int
	copy(sliceX, newSlice)
	start := time.Now()
	sort.Slice(sliceX, func(i, j int) bool { return sliceX[i] < sliceX[j] })
	end := time.Since(start)
	fmt.Printf("Slice took %s to sort using Slice()\n", end)

	start = time.Now()
	sort.SliceStable(newSlice, func(i, j int) bool { return newSlice[i] < newSlice[j] })
	end = time.Since(start)
	fmt.Printf("Slice took %s to sort using SliceStable()\n", end)

}
