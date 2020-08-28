package json_test

import (
	"testing"

	"github.com/xcaballero/contentLibrary-go/pkg/storage/json"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewStorage(t *testing.T) {
	s, err := json.NewStorage()
	require.NoError(t, err)
	assert.NotNil(t, s)
}
