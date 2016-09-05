package flagutil

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEnvOrDef(t *testing.T) {
	s := EnvOrDefault("foobar", "abcd")
	require.Equal(t, "abcd", s)

	os.Setenv("foobar", "efgh")
	s = EnvOrDefault("foobar", "abcd")
	require.Equal(t, "efgh", s)
}
