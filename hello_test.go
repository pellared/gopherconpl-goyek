package hello

import "testing"

func TestHi(t *testing.T) {
	want := "Hello World!"
	if got := Hi(); got != want {
		t.Errorf("Hi() = %v, want %v", got, want)
	}
}
