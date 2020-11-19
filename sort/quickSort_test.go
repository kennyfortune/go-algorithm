package sort

import (
	"fmt"
	"testing"
)

func Test_quickSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "base", args: args{arr: []int{5, 2, 3, 6, 1, 4}}},
		{name: "base", args: args{arr: []int{7, 5, 2, 8, 3, 6, 1, 9, 4}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			QuickSort(tt.args.arr)
			fmt.Println(tt.args.arr)
		})
	}
}
