package storage_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xcaballero/contentLibrary-go/pkg/storage"
)

func TestGetID(t *testing.T) {
	id, err := storage.GetID("testing")

	require.NoError(t, err)

	assert.True(t, strings.HasPrefix(id, "testing_"))
	assert.Len(t, id, 32)
}

func TestGetIDEmptyPrefix(t *testing.T) {
	id, err := storage.GetID("")

	require.NoError(t, err)

	assert.True(t, strings.HasPrefix(id, "_"))
	assert.Len(t, id, 25)
}
