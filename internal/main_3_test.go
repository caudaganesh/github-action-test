package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_doSomething3(t *testing.T) {
	t.Skip()
	t.Skip()
	t.Run("do something", func(t *testing.T) {
		got := doSomething()
		assert.Equal(t, "Hello, world.", got)
	})
}
