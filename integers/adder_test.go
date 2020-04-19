package integers

import "testing"

func TestAdder(t *testing.T) {
	actual := Add(2, 2)
	expected := 4

	if actual != expected {
		t.Errorf("expected: %d, but got %d", expected, actual)
	}
}
