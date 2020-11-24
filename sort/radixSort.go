package sort

func RadixSort(arr []int) {
	var save [10][]int

	base := 1
	for {
		//收集
		for i := range arr {
			tmp := arr[i] / base
			tmp %= base * 10
			save[tmp] = append(save[tmp], arr[i])
		}
		if len(save[0]) == len(arr) {
			break
		}

		//分发
		pos := 0
		for i := range save {
			if len(save[i]) != 0 {
				for _, j := range save[i] {
					arr[pos] = j
					pos++
				}
				save[i] = []int{}
			}
		}
		//循环变量
		base *= 10
	}
}
