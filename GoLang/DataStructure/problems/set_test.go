package main

import "testing"

func Test_distinctValues(t *testing.T) {
	
	tests := []struct {
		name string
		args []int
		want int
	}{
		{
			"neg",
			[]int{-1, -1, 0, 1, 1, 1 },
			2,
		},
		{
			"normal",
			[]int{ -2, -1, 0, 1, 2, 3 },
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distinctValues(tt.args); got != tt.want {
				t.Errorf("distinctValues() = %v, want %v", got, tt.want)
			}
		})
	}
}
