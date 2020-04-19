package main

import "testing"

func TestHello(t *testing.T) {
	tests := []struct {
		name string
		want string
		arg  string
	}{
		{
			name: "SpecificName",
			want: "Hello, Chris",
			arg:  "Chris",
		},
		{
			name: "DefaultBehavior",
			want: "Hello, World",
			arg:  "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hello(tt.arg); got != tt.want {
				t.Errorf("Hello() = %v, want %v", got, tt.want)
			}
		})
	}
}
