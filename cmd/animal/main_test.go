package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainCommandArgNumberError(t *testing.T) {
	t.Parallel()
	cmd := NewRootCmd()
	cmd.SetArgs([]string{"1arg"})
	err := cmd.Execute()
	assert.NotNil(t, err)
	assert.Contains(t, string(err.Error()), `accepts 2 arg(s), received 1`)

}

func TestMainCommandNormal(t *testing.T) {
	t.Parallel()
	cmd := NewRootCmd()
	cmd.SetArgs([]string{"2", "arg"})
	err := cmd.Execute()
	assert.Nil(t, err)
}
