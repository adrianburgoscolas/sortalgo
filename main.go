package main

import (
	"fmt"
	"time"
)

func sort(arr []int) {
	for i := 1; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			if arr[i-1] > arr[j] {
				arr[i-1], arr[j] = arr[j], arr[i-1]
			}
		}
	}
}

func countSort(arr []int) {
	var rang int
	for _, val := range arr {
		if rang < val {
			rang = val
		}
	}
	freqSlc := make([]int, rang+1)
	for _, val := range arr {
		freqSlc[val]++
	}
	offset := 0
	for num, amount := range freqSlc {
		if amount == 0 {
			continue
		}
		for i := offset; i < offset+amount; i++ {
			arr[i] = num
		}
		offset += amount
	}
}

func main() {
	startTime := time.Now()
	sort(SlcA)
	fmt.Printf("Enlapsed time: %s\n", time.Since(startTime))
	startTime = time.Now()
	cSort(SlcB)
	fmt.Printf("Enlapsed time: %s\n", time.Since(startTime))
	equal := true
	for i, val := range SlcA {
		if SlcB[i] != val {
			equal = false
		}
	}
	if equal {
		fmt.Println("Ok")
	}
}
