package randresources

import (
	"fmt"
	"testing"
)

func TestImage(t *testing.T) {
	img, err := BuildImage("test")
	if err != nil {
		t.Errorf("Error creating image")
	}
	fmt.Printf("%s\n", img.Base64())
}
