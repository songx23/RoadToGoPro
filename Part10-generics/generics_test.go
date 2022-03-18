package main

import "testing"

func TestSliceContains_Int(t *testing.T) {
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
			name: "int contains",
			args: args{
				src: []int{1, 2, 3},
				trg: 3,
			},
			want: true,
		},
		{
			name: "int does not contain",
			args: args{
				src: []int{1, 2, 3},
				trg: 4,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceContains(tt.args.src, tt.args.trg); got != tt.want {
				t.Errorf("SliceContains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceContains_String(t *testing.T) {
	type args struct {
		src []string
		trg string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "string contains",
			args: args{
				src: []string{"1", "2", "3"},
				trg: "3",
			},
			want: true,
		},
		{
			name: "string does not contain",
			args: args{
				src: []string{"1", "2", "3"},
				trg: "5",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceContains(tt.args.src, tt.args.trg); got != tt.want {
				t.Errorf("SliceContains() = %v, want %v", got, tt.want)
			}
		})
	}
}
