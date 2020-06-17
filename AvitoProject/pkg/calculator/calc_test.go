package calculator

import (
	"testing"
)

func TestCalculator_Calculate(t *testing.T) {
	a1:=2
	a2:=3
	exp := a1 + a2
	c := NewCalculator()
	res := c.Calculate(a1,a2)
	if res != exp{
		t.Errorf("got %v want %v",
			res, exp)
	}
}