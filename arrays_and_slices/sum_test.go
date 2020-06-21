package arrays_and_slices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1 + 2 + 3 + 4 + 5",
			args{[]int{1, 2, 3, 4, 5}},
			15,
		},
		{
			"1 + 2 + 3",
			args{[]int{1, 2, 3}},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.numbers); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumAll(t *testing.T) {
	type args struct {
		numbers1 []int
		numbers2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"base",
			args{
				[]int{1, 2},
				[]int{0, 9},
			},
			[]int{3, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumAll(tt.args.numbers1, tt.args.numbers2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumAllTails(t *testing.T) {
	type args struct {
		numbersToSum [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "make the sum of some slices",
			args: args{
				[][]int{
					{1, 2},
					{0, 9},
				},
			},
			want: []int{2, 9},
		},
		{
			name: "safely sum empty tests",
			args: args{
				[][]int{
					{},
					{3, 4, 5},
				},
			},
			want: []int{0, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumAllTails(tt.args.numbersToSum...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumAllTails() = %v, want %v", got, tt.want)
			}
		})
	}
}
