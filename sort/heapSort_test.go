package sort

import (
	"fmt"
	"testing"
)

func TestHeapSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "base", args: args{arr: []int{6, 2, 1, 5, 7, 9, 0, 3, 4, 8}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HeapSort(tt.args.arr)
			fmt.Println(tt.args.arr)
		})
	}
}
