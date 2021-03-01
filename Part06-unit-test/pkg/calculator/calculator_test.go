package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlus(t *testing.T) {
	res := Plus(1, 2)
	assert.Equal(t, 3.0, res)
	// native way to assert results
	res = Plus(1, -1)
	if res != 0.0 {
		t.Error("1 + (-1) should equal to 0.0")
	}
}

func TestMinus(t *testing.T) {
	res := Minus(1, 2)
	assert.Equal(t, -1.0, res)
	res = Minus(1, -1)
	assert.Equal(t, 2.0, res)
}

func TestTimes(t *testing.T) {
	res := Times(2, 3)
	assert.Equal(t, 6.0, res)
	res = Times(5, 0)
	assert.Equal(t, 0.0, res)
}

func TestDivide(t *testing.T) {
	res, err := Divide(5, 2)
	assert.NoError(t, err)
	assert.Equal(t, 2.5, res)
	res, err = Divide(5, 0)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "cannot divide by 0")
	assert.Equal(t, 0.0, res)
}

func TestTableDrivenTestPlus(t *testing.T) {
	type test struct {
		input1  float64
		input2  float64
		outcome float64
	}
	tests := []test{
		{input1: 1, input2: 2, outcome: 3},
		{input1: 1, input2: -1, outcome: 0},
	}
	for _, tc := range tests {
		actual := Plus(tc.input1, tc.input2)
		assert.Equal(t, tc.outcome, actual)
	}
}
