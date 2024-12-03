package day03

import "testing"

func Test_p1(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "sample",
			args: args{
				path: "inputs/sample.txt",
			},
			want: 161,
		},
		{
			name: "p1",
			args: args{
				path: "inputs/p1.txt",
			},
			want: 178794710,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p1(tt.args.path); got != tt.want {
				t.Errorf("p1() %v, want %v.", got, tt.want)
			}
		})
	}
}

func Test_p2(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "sample",
			args: args{
				path: "inputs/sample_p2.txt",
			},
			want: 48,
		},
		{
			name: "p2",
			args: args{
				path: "inputs/p1.txt",
			},
			want: 76729637,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p2(tt.args.path); got != tt.want {
				t.Errorf("p2() %v, want %v.", got, tt.want)
			}
		})
	}
}
