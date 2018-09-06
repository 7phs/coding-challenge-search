package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonArrayString_Unmarshal(t *testing.T) {
	exist := JsonArrayString(`["hello world","привет"]`).Unmarshal()
	expected := []string{"hello world", "привет"}

	assert.Equal(t, expected, exist)
}
