package main

import (
	"fmt"
	"math/rand"
	"time"
)

func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := i+1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

func main(){
	testArr := make([]int, 0, 10)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10; i++ {
		testArr = append(testArr, r.Intn(100))
	}
	fmt.Println("testArr", testArr)
	sortedArr := BubbleSort(testArr)
	fmt.Println("sortedArr", sortedArr)
}
