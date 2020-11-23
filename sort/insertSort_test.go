package sort

import (
	"fmt"
	"testing"
)

func TestInsertSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "base", args: args{arr: []int{5, 3, 6, 1, 2, 9, 7, 8, 0, 4}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InsertSort(tt.args.arr)
			fmt.Println(tt.args.arr)
		})
	}
}
