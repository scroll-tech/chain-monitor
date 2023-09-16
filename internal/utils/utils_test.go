package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetFuncName(t *testing.T) {
	name := GetFuncName()
	assert.Equal(t, "TestGetFuncName", name)
}
