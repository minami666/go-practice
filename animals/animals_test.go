package animals

import (
	"testing"
)

func TestSnakeFeed(t *testing.T) {
	expect := "Frog"
	actual := SnakeFeed()

	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}
}
