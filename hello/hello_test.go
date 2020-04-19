package main

import "testing"

func TestHello(t *testing.T) {
	tests := []struct {
		name string
		want string
		arg  string
		arg2 string
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
		{
			name: "InSpanish",
			want: "Hola, Elodie",
			arg:  "Elodie",
			arg2: "Spanish",
		},
		{
			name: "InFrench",
			want: "Bonjour, Luis",
			arg:  "Luis",
			arg2: "French",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hello(tt.arg, tt.arg2); got != tt.want {
				t.Errorf("Hello() = %v, want %v", got, tt.want)
			}
		})
	}
}
