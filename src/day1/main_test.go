package day01

import "testing"

func Test_p1(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "sample",
			args: args{
				path: "inputs/sample.txt",
			},
			want: 11,
		},
		{
			name: "p1",
			args: args{
				path: "inputs/p1.txt",
			},
			want: 1646452,
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
				path: "inputs/sample.txt",
			},
			want: 31,
		},
		{
			name: "p2",
			args: args{
				path: "inputs/p1.txt",
			},
			want: 23609874,
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
