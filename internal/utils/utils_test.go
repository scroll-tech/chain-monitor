package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFuncName(t *testing.T) {
	name := GetFuncName()
	assert.Equal(t, "TestGetFuncName", name)
}
