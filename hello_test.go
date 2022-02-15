package homework1

import (
	"fmt"
	"testing"
)

func Test_Hello(t *testing.T) {
	got := HelloWithEmoji()
	want := "Hello 🗺️ !"
	fmt.Println(want)
	if got != want {
		t.Errorf("got %q, want %q",got,want)
	}	
}
