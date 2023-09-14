package models

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidate(t *testing.T) {
	t.Parallel()

	t.Run("Pass validate", func(t *testing.T) {
		dataEmail := NewDataEmail("I", "to", "test-validate", "my-test-title", `<h1>test-send</h1>`, "aaaa,bbbb,cccc", "")

		err := dataEmail.Validate()
		require.NoError(t, err)
	})

	t.Run("Test title empty", func(t *testing.T) {
		dataEmail := NewDataEmail("", "to", "test-validate", "title", `<h1>test-send</h1>`, "", "")

		err := dataEmail.Validate()
		require.Error(t, err, ErrInvalidTo.Error)
	})

	t.Run("Pass validate and trim space", func(t *testing.T) {
		expcted := "NoSpace"
		dataEmail := NewDataEmail("   NoSpace       ", "to", "test-validate", "my-test-title", `<h1>test-send</h1>`, "", "")

		err := dataEmail.Validate()
		require.NoError(t, err)
		require.Equal(t, dataEmail.To, expcted)
	})
}
