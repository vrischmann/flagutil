package flagutil_test

import (
	"flag"
	"io"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/vrischmann/flagutil"
)

func mkFS() *flag.FlagSet {
	fs := flag.NewFlagSet("foobar", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	return fs
}

func TestNetworkAddress(t *testing.T) {
	fs := mkFS()

	var addr flagutil.NetworkAddress
	fs.Var(&addr, "h", "Address")

	s := []string{"-h", "a:4000"}

	err := fs.Parse(s)
	require.Nil(t, err)
	require.Equal(t, flagutil.NetworkAddress("a:4000"), addr)

	s = []string{"-h", "foo"}
	err = fs.Parse(s)
	require.NotNil(t, err)
}

func TestNetworkAddresses(t *testing.T) {
	fs := mkFS()

	var addrs flagutil.NetworkAddresses
	fs.Var(&addrs, "H", "Addresses")

	s := []string{"-H", "a:4000,b:5000"}

	err := fs.Parse(s)
	require.Nil(t, err)
	require.Equal(t, 2, len(addrs))
	require.Equal(t, "a:4000", addrs[0])
	require.Equal(t, "b:5000", addrs[1])
	require.Equal(t, "a:4000,b:5000", addrs.String())
	require.Equal(t, []string{"a:4000", "b:5000"}, addrs.StringSlice())

	s = []string{"-H", "foo,bar"}
	err = fs.Parse(s)
	require.NotNil(t, err)
}

func TestStrings(t *testing.T) {
	fs := mkFS()

	var strings flagutil.Strings
	fs.Var(&strings, "s", "Strings")

	s := []string{"-s", "foo,bar,baz"}

	err := fs.Parse(s)
	require.Nil(t, err)
	require.Equal(t, 3, len(strings))
	require.Equal(t, "foo", strings[0])
	require.Equal(t, "bar", strings[1])
	require.Equal(t, "baz", strings[2])
	require.Equal(t, "foo,bar,baz", strings.String())
}

func TestURL(t *testing.T) {
	fs := mkFS()

	var url flagutil.URL
	fs.Var(&url, "u", "URL")

	s := []string{"-u", "https://google.com"}

	err := fs.Parse(s)
	require.Nil(t, err)
	require.True(t, url.IsValid())
	require.Equal(t, "google.com", url.URL.Host)
	require.Equal(t, "https", url.URL.Scheme)
	require.Equal(t, "https://google.com", url.String())

	s = []string{"-u", "://foobar"}
	err = fs.Parse(s)
	require.NotNil(t, err)
}

func TestURLs(t *testing.T) {
	fs := mkFS()

	var urls flagutil.URLs
	fs.Var(&urls, "U", "URLs")

	s := []string{"-U", "https://google.com,https://google.de"}

	err := fs.Parse(s)
	require.Nil(t, err)
	require.Equal(t, 2, len(urls))
	require.Equal(t, "https://google.com", urls[0].String())
	require.Equal(t, "https://google.de", urls[1].String())
	require.Equal(t, "https://google.com,https://google.de", urls.String())

	s = []string{"-U", "://foobar"}
	err = fs.Parse(s)
	require.NotNil(t, err)
}

func TestDuration(t *testing.T) {
	fs := mkFS()

	var d flagutil.Duration
	fs.Var(&d, "d", "Duration")

	testCases := []struct {
		in  string
		exp time.Duration
	}{
		{"10s", 10 * time.Second},
		{"10ms", 10 * time.Millisecond},
		{"10m", 10 * time.Minute},
		{"10h", 10 * time.Hour},
	}

	for _, tc := range testCases {
		err := fs.Parse([]string{"-d", tc.in})
		require.Nil(t, err)

		require.Equal(t, tc.exp, d.Duration)
	}
}
