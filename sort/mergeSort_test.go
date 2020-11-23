package sort

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "base", args: args{arr: []int{5, 2, 1, 8, 4, 9, 0, 3, 7, 6}}, want: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(tt.args.arr)
			if got := MergeSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSort() = %v, want %v", got, tt.want)
			} else {
				fmt.Println(got)
			}
		})
	}
}

func TestMergeSortIteration(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{

		{name: "base", args: args{arr: []int{5, 2, 1, 8, 4, 9, 0, 3, 7, 6}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MergeSortIteration(tt.args.arr)
			fmt.Println(tt.args.arr)
		})
	}
}
