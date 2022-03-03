package main

import "testing"

func Test_sumUpWithLock(t *testing.T) {
	type args struct {
		start int
		end   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Sum up correctly",
			args: args{
				start: 1,
				end:   100,
			},
			want: 5050,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumUpWithLock(tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("sumUpWithLock() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sumUpWithRWLock(t *testing.T) {
	type args struct {
		start int
		end   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Sum up and inspect correctly",
			args: args{
				start: 1,
				end:   100,
			},
			want: 5050,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumUpWithRWLock(tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("sumUpWithRWLock() = %v, want %v", got, tt.want)
			}
		})
	}
}
