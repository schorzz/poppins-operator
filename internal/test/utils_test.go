package test

import (
	"github.com/schorzz/poppins-operator/internal/pkg/api/rest"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUtils(t *testing.T)  {
	list := []string{"default", "13"}
	elem := "default"
	assert.True(t, rest.ListContains(list, elem))
	assert.False(t, rest.ListContains(list, "not in there"))
}