package sort

import (
	"fmt"
	"testing"
)

func TestCountingSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "base", args: args{arr: []int{2, 6, 4, 8, 0, 8, 5, 7, 3, 9}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CountingSort(tt.args.arr)
			fmt.Println(tt.args.arr)

		})
	}
}
