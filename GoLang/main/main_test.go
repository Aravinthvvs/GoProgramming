package main

import "testing"

func Test_getLongestCommonPrefix(t *testing.T) {
	inputStrcut := []string{"aravinth", "arav", "arav"}
	type args struct {
		str []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"normal",
		args{inputStrcut},
		"arav",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLongestCommonPrefix(tt.args.str); got != tt.want {
				t.Errorf("getLongestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
