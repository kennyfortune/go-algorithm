package sort

func ShellSort(arr []int)  {
	for step := len(arr) / 2; step >= 1; step /= 2{
		for i := step; i < len(arr); i++ {
			//满足条件则插入
			for j := i - step; j >= 0 && arr[j+step] < arr[j]; j -= step {
				arr[j], arr[j+step] = arr[j+step], arr[j]
			}
		}
	}

}
