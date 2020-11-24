package sort

import (
	"fmt"
	"testing"
)

func TestRadixSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "base", args: args{arr: []int{11, 27, 16, 35, 19, 55, 22, 27, 40, 31, 38, 44, 58, 45, 57, 46, 60}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(tt.args.arr)
			RadixSort(tt.args.arr)
			fmt.Println(tt.args.arr)
		})
	}
}
