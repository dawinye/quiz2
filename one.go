package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

//https://go.dev/tour/concurrency/2
func sumFunc(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}
func main() {
	var intArg, _ = strconv.Atoi(os.Args[1])
	var sliceX []int
	var sum int
	for i := 0; i < intArg; i++ {
		sliceX = append(sliceX, rand.Intn(intArg))
	}
	start := time.Now()
	for _, number := range sliceX {
		sum += number
	}
	end := time.Since(start)
	fmt.Printf("Sum is %d \n Slice took %s to sum up\n", sum, end)

	start = time.Now()
	c := make(chan int)
	go sumFunc(sliceX[:len(sliceX)/2], c)
	go sumFunc(sliceX[len(sliceX)/2:], c)
	x, y := <-c, <-c // receive from c
	end = time.Since(start)
	fmt.Printf("Sum is %d \nSlice took %s to sum up\n", x+y, end)
}
