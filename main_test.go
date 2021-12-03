package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_doSomething(t *testing.T) {
	t.Run("do something", func(t *testing.T) {
		got := doSomething()
		assert.Equal(t, "Hello, world.", got)
	})
}
