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

	t.Run("single", func(t *testing.T) {
		t.Run("valid", func(t *testing.T) {
			var addr flagutil.NetworkAddress
			fs.Var(&addr, "h", "Address")

			s := []string{"-h", "a:4000"}

			err := fs.Parse(s)
			require.NoError(t, err)
			require.Equal(t, flagutil.NetworkAddress("a:4000"), addr)
		})

		t.Run("invalid", func(t *testing.T) {
			s := []string{"-h", "foo"}
			err := fs.Parse(s)
			require.EqualError(t, err, `invalid value "foo" for flag -h: address foo: missing port in address`)
		})
	})
	t.Run("multiple", func(t *testing.T) {
		t.Run("valid", func(t *testing.T) {
			var addrs flagutil.NetworkAddresses
			fs.Var(&addrs, "H", "Addresses")

			s := []string{"-H", "a:4000,b:5000"}

			err := fs.Parse(s)
			require.NoError(t, err)
			require.Equal(t, 2, len(addrs))
			require.Equal(t, "a:4000", addrs[0])
			require.Equal(t, "b:5000", addrs[1])
			require.Equal(t, "a:4000,b:5000", addrs.String())
			require.Equal(t, []string{"a:4000", "b:5000"}, addrs.StringSlice())
		})

		t.Run("invalid", func(t *testing.T) {
			s := []string{"-H", "foo,bar"}
			err := fs.Parse(s)
			require.EqualError(t, err, `invalid value "foo,bar" for flag -H: address foo: missing port in address`)
		})
	})

}

func TestStrings(t *testing.T) {
	fs := mkFS()

	t.Run("valid", func(t *testing.T) {
		var strings flagutil.Strings
		fs.Var(&strings, "s", "Strings")

		s := []string{"-s", "foo,bar,baz"}

		err := fs.Parse(s)
		require.NoError(t, err)
		require.Equal(t, 3, len(strings))
		require.Equal(t, "foo", strings[0])
		require.Equal(t, "bar", strings[1])
		require.Equal(t, "baz", strings[2])
		require.Equal(t, "foo,bar,baz", strings.String())
	})
}

func TestURL(t *testing.T) {
	fs := mkFS()

	t.Run("one", func(t *testing.T) {
		t.Run("valid", func(t *testing.T) {
			var url flagutil.URL
			fs.Var(&url, "u", "URL")

			s := []string{"-u", "https://google.com"}

			err := fs.Parse(s)
			require.NoError(t, err)
			require.True(t, url.IsValid())
			require.Equal(t, "google.com", url.Host)
			require.Equal(t, "https", url.Scheme)
			require.Equal(t, "https://google.com", url.String())
		})

		t.Run("invalid", func(t *testing.T) {
			s := []string{"-u", "://foobar"}
			err := fs.Parse(s)
			require.EqualError(t, err, `invalid value "://foobar" for flag -u: parse "://foobar": missing protocol scheme`)
		})
	})

	t.Run("multiple", func(t *testing.T) {
		t.Run("valid", func(t *testing.T) {
			var urls flagutil.URLs
			fs.Var(&urls, "U", "URLs")

			s := []string{"-U", "https://google.com,https://google.de"}

			err := fs.Parse(s)
			require.NoError(t, err)
			require.Equal(t, 2, len(urls))
			require.Equal(t, "https://google.com", urls[0].String())
			require.Equal(t, "https://google.de", urls[1].String())
			require.Equal(t, "https://google.com,https://google.de", urls.String())
		})
		t.Run("invalid", func(t *testing.T) {
			s := []string{"-U", "://foobar,://lol"}
			err := fs.Parse(s)
			require.EqualError(t, err, `invalid value "://foobar,://lol" for flag -U: parse "://foobar": missing protocol scheme`)
		})
	})

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
		t.Run("", func(t *testing.T) {
			err := fs.Parse([]string{"-d", tc.in})
			require.NoError(t, err)

			require.Equal(t, tc.exp, d.Duration)
		})
	}
}
