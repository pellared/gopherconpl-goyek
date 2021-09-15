package hello

import "testing"

func TestHi(t *testing.T) {
	want := "Hello world"
	if got := Hi(); got != want {
		t.Errorf("Hi() = %v, want %v", got, want)
	}
}
