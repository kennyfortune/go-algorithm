package sort

import (
	"fmt"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		{name:"base",args: args{arr: []int{3,5,6,2,1,7,4}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SelectionSort(tt.args.arr)
			fmt.Println(tt.args.arr)
		})
	}
}
