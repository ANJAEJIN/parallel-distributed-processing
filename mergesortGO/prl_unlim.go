package main

import (
	"fmt"
	"io/ioutil"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func main() {
	dat, err := ioutil.ReadFile("./random.txt")
	check(err)
	slice := make([]int, 1000000)
	s := strings.Split(string(dat), "\n")
	for i := 0; i < len(s); i++ {
		slice[i], err = strconv.Atoi(s[i])
		check(err)
	}

	runtime.GOMAXPROCS(runtime.NumCPU()) // CPU 개수를 구한 뒤 사용할 최대 CPU 개수 설정

	fmt.Println(runtime.GOMAXPROCS(0)) // 설정 값 출력

	startTime := time.Now()
	result := make(chan []int)
	go MergeSort(slice, result)
	r := <-result
	elapsedTime := time.Since(startTime)
	fmt.Printf("run time: %s\n", elapsedTime)

	fmt.Println("\n--- Unsorted --- \n\n", slice[0:10])
	fmt.Println("\n--- Sorted ---\n\n", r[0:10], "\n")
}

func Merge(ldata []int, rdata []int) (result []int) {
	result = make([]int, len(ldata)+len(rdata))
	lidx, ridx := 0, 0

	for i := 0; i < cap(result); i++ {
		switch {
		case lidx >= len(ldata):
			result[i] = rdata[ridx]
			ridx++
		case ridx >= len(rdata):
			result[i] = ldata[lidx]
			lidx++
		case ldata[lidx] < rdata[ridx]:
			result[i] = ldata[lidx]
			lidx++
		default:
			result[i] = rdata[ridx]
			ridx++
		}
	}

	return
}

func MergeSort(data []int, r chan []int) {
	if len(data) == 1 {
		r <- data
		return
	}

	leftChan := make(chan []int)
	rightChan := make(chan []int)
	middle := len(data) / 2

	go MergeSort(data[:middle], leftChan)
	go MergeSort(data[middle:], rightChan)

	/*
		MergeSort(data[:middle], leftChan)
		MergeSort(data[middle:], rightChan)
	*/

	ldata := <-leftChan
	rdata := <-rightChan

	close(leftChan)
	close(rightChan)
	r <- Merge(ldata, rdata)

	return
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
