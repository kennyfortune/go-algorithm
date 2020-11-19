package sort

func QuickSort(arr []int) {

	if len(arr) <= 1 {
		return
	}
	start := 0
	end := len(arr) - 1
	for start<end {
		for start<end {
			if arr[start] <= arr[end] {
				start++
			} else {
				tmp := arr[start]
				arr[start] = arr[end]
				arr[end] = tmp
				break
			}
		}
		for start < end {
			if arr[start] <= arr[end]{
				end--
			} else {
				arr[start], arr[end] = arr[end] , arr[start]
				break
			}
		}
	}
	defer QuickSort(arr[0:start])
	defer QuickSort(arr[start+1:])
}
