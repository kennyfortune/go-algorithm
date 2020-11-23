package sort

func HeapSort(arr []int) {
	Adjust(arr, len(arr))
	for i := len(arr) - 1; i >= 1; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		Down(arr[:i], 0)
	}

}

func Adjust(arr []int, len int) {
	for i := len/2 - 1; i >= 0; i-- {
		Down(arr, i)
	}
}

func Down(arr []int, pos int) {
	if pos*2+1 >= len(arr) {
		return
	}
	if arr[pos] < arr[pos*2+1] {
		arr[pos], arr[pos*2+1] = arr[pos*2+1], arr[pos]
		Down(arr, pos*2+1)
	}
	if pos*2+2 >= len(arr) {
		return
	}
	if arr[pos] < arr[pos*2+2] {
		arr[pos], arr[pos*2+2] = arr[pos*2+2], arr[pos]
		Down(arr, pos*2+2)
	}
}
