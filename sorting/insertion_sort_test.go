package sorting

import (
	"reflect"
	"testing"
)

func TestInsertionSortInt(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Insertion sort 1",
			args: args{
				arr: []int{77, 99, 44, 55, 22, 88, 11, 0, 66, 33},
			},
			want: []int{0, 11, 22, 33, 44, 55, 66, 77, 88, 99},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertionSortInt(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertionSortInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
