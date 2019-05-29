package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func main() {
	dat, err := ioutil.ReadFile("./random.txt")
	check(err)
	slice := make([]int, 1000000)
	s := strings.Split(string(dat),"\n")
	for i:= 0;i < len(s)-1 ; i++ {
		slice[i],err = strconv.Atoi(s[i])
		check(err)
	}


	startTime := time.Now()
	result := make(chan []int)
	go MergeSort(slice, result)
	r := <- result
	elapsedTime := time.Since(startTime)
	fmt.Printf("run time: %s\n", elapsedTime)


	fmt.Println("\n--- Unsorted --- \n\n", slice[0:10])
	fmt.Println("\n--- Sorted ---\n\n", r[0:100], "\n")
}

func Merge(ldata []int, rdata []int) (result []int) {
	result = make([]int, len(ldata) + len(rdata))
	lidx, ridx := 0, 0

	for i:=0;i<cap(result);i++ {
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

	if len(data)<=1000 {
		r <- mergeSort(data)
	}else {
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
	}



	return
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}


func mergeSort(items []int) []int {
	var num = len(items)

	if num == 1 {
		return items
	}

	middle := int(num / 2)
	var (
		left = make([]int, middle)
		right = make([]int, num-middle)
	)
	for i := 0; i < num; i++ {
		if i < middle {
			left[i] = items[i]
		} else {
			right[i-middle] = items[i]
		}
	}

	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []int) (result []int) {
	result = make([]int, len(left) + len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return
}