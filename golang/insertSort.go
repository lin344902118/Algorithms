package golang

func InsertSort(arr []int) []int {
	new_arr := make([]int, 0, len(arr))
	new_arr = append(new_arr, arr[0])
	for i := 1; i < len(arr); i++ {
		inserted := false
		for j := 0; j < len(new_arr); j++ {
			if arr[i] < new_arr[j] {
				tmp_arr := make([]int, len(new_arr[j:]))
				copy(tmp_arr, new_arr[j:])
				new_arr = append(new_arr[:j], arr[i])
				new_arr = append(new_arr, tmp_arr...)
				inserted = true
				break
			}
		}
		if !inserted {
			new_arr = append(new_arr, arr[i])
		}
	}
	return new_arr
}

