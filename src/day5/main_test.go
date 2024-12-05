package day05

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
			want: 143,
		},
		{
			name: "p1",
			args: args{
				path: "inputs/p1.txt",
			},
			want: 6242,
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
			want: 123,
		},
		{
			name: "p2",
			args: args{
				path: "inputs/p1.txt",
			},
			want: 5169,
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
