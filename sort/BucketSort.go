package sort

func BucketSort(arr []int) {
	var min, max int
	for _, v := range arr {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	//桶初始化
	defaultSize := 5
	var buckets [5][]int
	arraySize := (max-min)/defaultSize + 1

	//分桶函数 这个函数决定了效率 可以有各种方式分桶
	for _, v := range arr {
		bucketPos := (v - min) / arraySize
		if buckets[bucketPos] == nil {
			buckets[bucketPos] = make([]int, 0)
		}
		buckets[bucketPos] = append(buckets[bucketPos], v)
	}

	p := 0
	//每个桶再排序 子排序可以用各种排序算法，也可以继续使用桶排序
	for i := range buckets {
		if len(buckets[i]) != 0 {
			InsertSort(buckets[i])
			for _, v := range buckets[i] {
				arr[p] = v
				p++
			}
		}
	}

}
