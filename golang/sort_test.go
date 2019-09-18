package golang

import (
	"fmt"
	"math/rand"
	"testing"
)

func GetRandomArray() (array []int) {
	array = make([]int, 0, 10)
	for i := 0; i < 10; i++ {
		randomNum := rand.Intn(100)
		array = append(array, randomNum)
	}
	return array
}

func SortTest(array []int) bool {
	for i := 0; i < len(array)-1; i++ {
		for j := i+1; j < len(array); j++ {
			if array[j] < array[i] {
				return false
			}
		}
	}
	return true
}

func TestBubbleSort(t *testing.T) {
	array := GetRandomArray()
	fmt.Println("random array", array)
	sortedArr := BubbleSort(array)
	fmt.Println("sorted array", sortedArr)
	if !SortTest(sortedArr) {
		t.Errorf("bubble sort err, before sort:%v, after sort:%v", array, sortedArr)
	}
}

func TestInsertSort(t *testing.T) {
	array := GetRandomArray()
	fmt.Println("random array", array)
	sortedArr := InsertSort(array)
	fmt.Println("sorted array", sortedArr)
	if !SortTest(sortedArr) {
		t.Errorf("bubble sort err, before sort:%v, after sort:%v", array, sortedArr)
	}
}

func TestMergeSort(t *testing.T) {
	array := GetRandomArray()
	fmt.Println("random array", array)
	sortedArr := MergeSort(array)
	fmt.Println("sorted array", sortedArr)
	if !SortTest(sortedArr) {
		t.Errorf("bubble sort err, before sort:%v, after sort:%v", array, sortedArr)
	}
}

func TestQuickSort(t *testing.T) {
	array := GetRandomArray()
	fmt.Println("random array", array)
	sortedArr := QuickSort(array)
	fmt.Println("sorted array", sortedArr)
	if !SortTest(sortedArr) {
		t.Errorf("bubble sort err, before sort:%v, after sort:%v", array, sortedArr)
	}
}

func TestSelectSort(t *testing.T) {
	array := GetRandomArray()
	fmt.Println("random array", array)
	sortedArr := SelectSort(array)
	fmt.Println("sorted array", sortedArr)
	if !SortTest(sortedArr) {
		t.Errorf("bubble sort err, before sort:%v, after sort:%v", array, sortedArr)
	}
}