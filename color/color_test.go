package color

import (
	"testing"

	"github.com/fatih/color"
)

func TestColor(t *testing.T) {
	n, err := color.New(color.FgHiBlue).Println("color")
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(n)
}
