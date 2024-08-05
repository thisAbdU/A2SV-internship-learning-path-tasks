package main

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

func main() {
	fmt.Println(Calculate(5))
}

func Calculate(n int) int {
	return n * n
}

func TestCalculate(t *testing.T) {
	assert.Equal(t, 25, Calculate(5))
}

func TestBulk(t *testing.T){

	assert := assert.New(t)

	var tests = []struct {
		input int
		expected int
	}{
		{5, 25},
		{0, 0},
		{1, 1},
		{-5, 25},
	}

	for _, test := range tests {
		assert.Equal(Calculate(test.input), test.expected)
	}
}

