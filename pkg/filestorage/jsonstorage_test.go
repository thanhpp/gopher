package filestorage_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thanhpp/gopher/pkg/filestorage"
)

func TestJSONStorage(t *testing.T) {
	var (
		path = "./test_json_storage.json"
	)
	defer os.RemoveAll(path)

	s, err := filestorage.NewJSONStorage(path)
	require.NoError(t, err)

	_, err = os.Open(path)
	require.NoError(t, err)

	err = s.Set("key", 10)
	require.NoError(t, err)

	v, err := s.Get("key")
	require.NoError(t, err)

	iV, ok := v.(int)
	require.True(t, ok)

	assert.Equal(t, iV, 10)

	assert.NoError(t, s.Del("key"))
}
