package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSortIteration(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "base", args: args{[]int{5, 3, 2, 4, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BubbleSort(tt.args.arr)
			fmt.Println(tt.args)
		})
	}
}
