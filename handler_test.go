package main

import "testing"

func Test_handleCount(t *testing.T) {
	type args struct {
		i int
	}
	var tests []struct {
		name string
		args args
		want int
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handleCount(tt.args.i); got != tt.want {
				t.Errorf("handleCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
