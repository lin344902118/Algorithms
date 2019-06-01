package main

import (
	"fmt"
	"math/rand"
	"time"
)

func FindSmallestNum(arr []int) int {
	smallestNum := arr[0]
	smallestNumIndex := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] < smallestNum {
			smallestNum = arr[i]
			smallestNumIndex = i
		}
	}
	return smallestNumIndex
}

func SelectSort(arr []int) []int{
	newArr := make([]int, 0, len(arr))
	for len(arr) > 0 {
		smallestNumIndex := FindSmallestNum(arr)
		newArr = append(newArr, arr[smallestNumIndex])
		arr = append(arr[:smallestNumIndex], arr[smallestNumIndex+1:]...)
	}
	return newArr
}

func main(){
	testArr := make([]int, 0, 10)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10; i++ {
		testArr = append(testArr, r.Intn(100))
	}
	fmt.Println("testArr", testArr)
	sortedArr := SelectSort(testArr)
	fmt.Println("sortedArr", sortedArr)
}

