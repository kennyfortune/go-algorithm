package sort

func MergeSort(arr []int) []int {
	i := len(arr) / 2
	if i == 0 {
		return arr
	}
	left := MergeSort(arr[:i])
	right := MergeSort(arr[i:])
	return Merge(left, right)

}

/**
 * @Description: 迭代版实现
 * @param arr
 * @return []int
 */
func MergeSortIteration(arr []int) {
	for i := 1; i < len(arr); i *= 2 {
		aleft := 0
		aright, bleft := i, i
		bright := 2 * i
		if bright > len(arr) {
			bright = len(arr)
		}
		for aright < len(arr) {
			MergeIteration(arr, aleft, aright, bleft, bright)
			aleft = bright
			aright = aleft + i
			bleft = aright
			bright = bleft + i
			if bright > len(arr) {
				bright = len(arr)
			}

		}
	}
}

func Merge(a, b []int) []int {
	result := make([]int, len(a)+len(b))
	i := 0
	j := 0
	for i != len(a) && j != len(b) {
		if a[i] <= b[j] {
			result[i+j] = a[i]
			i++
		} else {
			result[i+j] = b[j]
			j++
		}
	}
	for i < len(a) {
		result[i+j] = a[i]
		i++
	}
	for j < len(b) {
		result[i+j] = b[j]
		j++

	}
	return result
}

func MergeIteration(arr []int, aleft, aright, bleft, bright int) {
	newa := make([]int, aright-aleft)
	newb := make([]int, bright-bleft)
	copy(newa, arr[aleft:aright])
	copy(newb, arr[bleft:bright])
	i := 0
	j := 0
	for i != len(newa) && j != len(newb) {
		if newa[i] <= newb[j] {
			arr[aleft+i+j] = newa[i]
			i++
		} else {
			arr[aleft+i+j] = newb[j]
			j++
		}
	}
	for i < len(newa) {
		arr[aleft+i+j] = newa[i]
		i++
	}
	for j < len(newb) {
		arr[aleft+i+j] = newb[j]
		j++

	}
}
