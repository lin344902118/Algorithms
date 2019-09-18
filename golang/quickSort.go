package golang

func QuickSort(arr []int) []int {
	if len(arr) < 2{
		return arr
	}
	num := arr[0]
	less := make([]int, 0)
	greater := make([]int, 0)
	sortedArr := make([]int, 0)
	for i := 1; i < len(arr); i++ {
		if arr[i] <= num {
			less = append(less, arr[i])
		} else {
			greater = append(greater, arr[i])
		}
	}
	sortedArr = append(QuickSort(less), num)
	sortedArr = append(sortedArr, QuickSort(greater)...)
	return sortedArr
}
