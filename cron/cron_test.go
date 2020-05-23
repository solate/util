package cron

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValid(t *testing.T) {

	testCases := []struct {
		cron     string
		expected bool
	}{
		{cron: "* * * * * ?", expected: true},
		{cron: "0 * * ? * * *", expected: true},
		{cron: "0 0/1 * * * ? *", expected: true},
		{cron: "* * * * * * *", expected: false},
		{cron: "* * * * * ? 1980", expected: false},
	}

	for _, test := range testCases {
		actual := Valid(test.cron)
		assert.Equal(t, test.expected, actual)
	}

}
