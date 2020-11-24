package sort

func CountingSort(arr []int) {
	var min, max int = arr[0], arr[0]
	for i := range arr {
		if min > arr[i] {
			min = arr[i]
		}
		if max < arr[i] {
			max = arr[i]
		}
	}
	save := make([]int, max-min+1)
	for _, v := range arr {
		save[v-min]++
	}
	pos := 0
	var v int
	for i := range arr {
		v, pos = FindNext(save, pos)
		arr[i] = v + min
	}
}

func FindNext(arr []int, pos int) (v int, newPos int) {
	for arr[pos] == 0 {
		pos++
	}
	arr[pos]--
	return pos, pos

}
