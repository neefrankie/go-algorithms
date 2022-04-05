package exercises

import (
	"reflect"
	"testing"
)

func TestSumList(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Sum int",
			args: args{
				arr: []int{8, 4, 2, 10, 19},
			},
			want: 43,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumList(tt.args.arr); got != tt.want {
				t.Errorf("SumList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLargestDifference(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Largest difference",
			args: args{
				arr: []int{1, 4, 9, 3, 10},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LargestDifference(tt.args.arr); got != tt.want {
				t.Errorf("LargestDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindDuplicates(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Find duplicates",
			args: args{
				arr: []int{2, 2, 3, 5, 7, 7, 7, 9, 10, 11, 12, 12, 13},
			},
			want: []int{2, 7, 12},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindDuplicates(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
