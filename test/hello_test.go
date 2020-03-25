package test

import (
	"github.com/ArseniySavin/IINBINGoCheck"
	"testing"
)

func TestExampleHello(t *testing.T) {
	IINBINGoCheck.ExampleHello()
	got := -1
	t.Errorf("Abs(-1) = %d; want 1", got)

}
