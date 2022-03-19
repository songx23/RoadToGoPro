package main

import "testing"

func TestIntSliceContains(t *testing.T) {
	type args struct {
		src []int
		trg int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "contains",
			args: args{
				src: []int{1, 2, 3},
				trg: 2,
			},
			want: true,
		},
		{
			name: "does not contain",
			args: args{
				src: []int{1, 2, 3},
				trg: 4,
			},
			want: false,
		},
		{
			name: "empty slice",
			args: args{
				src: []int{},
				trg: 5,
			},
			want: false,
		},
		{
			name: "nil slice",
			args: args{
				src: nil,
				trg: 5,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntSliceContains(tt.args.src, tt.args.trg); got != tt.want {
				t.Errorf("IntSliceContains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkIntSliceContains(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntSliceContains([]int{1, 2, 3}, 3)
	}
}
