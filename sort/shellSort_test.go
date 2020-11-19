package sort

import (
	"fmt"
	"testing"
)

func TestShellSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		{name:"base", args: args{arr: []int{5,1,7,2,4,3,6}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ShellSort(tt.args.arr)
			fmt.Println(tt.args.arr)
		})
	}
}
