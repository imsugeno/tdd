package money

import "testing"

func TestMultiplication(t *testing.T) {
	var five Dollar
	five.amount = 10
	five.times(2)
	if five.amount != 10 {
		t.Error()
	}
}
