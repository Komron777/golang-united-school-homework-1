package solution

import (
	"fmt"
	"testing"
)

func Test_Hello(t *testing.T) {
	got := GetMessage()
	want := "Hello 🗺️ !"
	fmt.Println(want)
	if got != want {
		t.Failed()
		//t.Errorf("got %q, want %q",got,want)
	}	
}
	