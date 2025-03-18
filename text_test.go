package randresources

import (
	"fmt"
	"testing"
)

func TestText(t *testing.T) {
	for i := 0; i < 100; i++ {
		nickname := Text()
		if nickname == "" {
			t.Errorf("Generated nickname is empty")
		}
		fmt.Printf("%s\n", nickname)
	}
}
