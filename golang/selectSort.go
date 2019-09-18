package golang

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
