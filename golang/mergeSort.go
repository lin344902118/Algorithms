package golang

func merge(arr1, arr2 []int) []int {
	tmp := make([]int, 0, len(arr1)+len(arr2))
	i, j := 0, 0
	for (i < len(arr1) && j < len(arr2)) {
		if arr1[i] < arr2[j] {
			tmp = append(tmp, arr1[i])
			i++
		} else {
			tmp = append(tmp, arr2[j])
			j++
		}
	}
	if i < len(arr1) {
		tmp = append(tmp, arr1[i:]...)
	}
	if j < len(arr2) {
		tmp = append(tmp, arr2[j:]...)
	}
	return tmp
}

func MergeSort(arr []int)[]int {
	if len(arr) <= 2{
		if len(arr) == 1 {
			return arr
		}
		if arr[0] > arr[1] {
			arr[0], arr[1] = arr[1], arr[0]
		}
		return arr
	}
	mid := len(arr) / 2
	if len(arr) % 2 != 0 {
		mid++
	}
	font := MergeSort(arr[:mid])
	after := MergeSort(arr[mid:])
	mergedArr := merge(font, after)
	return mergedArr
}
