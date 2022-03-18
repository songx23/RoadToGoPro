package main

import "testing"

func TestStringSliceContains(t *testing.T) {
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
			name: "contains",
			args: args{
				src: []string{"a", "b", "c"},
				trg: "b",
			},
			want: true,
		},
		{
			name: "does not contain",
			args: args{
				src: []string{"a", "b", "c"},
				trg: "d",
			},
			want: false,
		},
		{
			name: "empty slice",
			args: args{
				src: []string{},
				trg: "b",
			},
			want: false,
		},
		{
			name: "nil slice",
			args: args{
				src: nil,
				trg: "b",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringSliceContains(tt.args.src, tt.args.trg); got != tt.want {
				t.Errorf("StringSliceContains() = %v, want %v", got, tt.want)
			}
		})
	}
}
