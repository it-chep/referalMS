package postgres

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestName(t *testing.T) {
	t.Parallel()
	t.Run("test of tests", func(t *testing.T) {
		t.Parallel()
		require.False(t, false)
	})
}
