package sort

func SelectionSort(arr []int){
	for i, _ := range arr {
		min := i
		for j:=i+1;j<len(arr);j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		arr[min], arr[i] = arr[i], arr[min]
	}
}
