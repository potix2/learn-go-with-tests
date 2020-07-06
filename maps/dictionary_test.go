package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	type args struct {
		dictionary Dictionary
		word string
	}
	tests := []struct {
		name string
		args args
		want string
		wantErr error
	}{
		{
			"known word",
			args{
				Dictionary{"test": "this is just a test"},
				"test",
			},
			"this is just a test",
			nil,
		},
		{
			"unknown word",
			args{
				Dictionary{"test": "this is just a test"},
				"unknown",
			},
			"",
			ErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.args.dictionary.Search(tt.args.word)
			if tt.wantErr != nil {
				if err != tt.wantErr {
					t.Fatalf("expected to get an error. but got %q", err)
				}
			}
			assertStrings(t, got, tt.want)
		})
	}
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
