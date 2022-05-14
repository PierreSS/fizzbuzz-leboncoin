package handlers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz(t *testing.T) {
	cases := []struct {
		inputInt1  int
		inputInt2  int
		inputLimit int
		inputStr1  string
		inputStr2  string
		output     string
	}{
		{3, 5, 16, "fizz", "buzz", "1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16"},
		{2, 3, 20, "lebon", "coin", "1,lebon,coin,lebon,5,leboncoin,7,lebon,coin,lebon,11,leboncoin,13,lebon,coin,lebon,17,leboncoin,19,lebon"},
	}

	for _, test := range cases {
		output := ExecFizzBuzz(test.inputInt1, test.inputInt2, test.inputLimit, test.inputStr1, test.inputStr2)
		assert.Equal(t, test.output, output)
	}
}
