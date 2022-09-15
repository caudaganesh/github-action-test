package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_doSomethingAgain(t *testing.T) {
	t.Run("do something", func(t *testing.T) {
		got := doSomethingAgain()
		assert.Equal(t, "Hello, world. Hello, world.", got)
	})
}
