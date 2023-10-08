package testing

import "testing"

func FooTest(t *testing.T) {
	if actualResult := foo(5, 5); actualResult != 10 {
		t.Errorf("expected %d, got %d", actualResult, 10)
	}
}
