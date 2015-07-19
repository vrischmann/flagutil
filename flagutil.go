package flagutil

import (
	"net"
	"strings"
)

// ListenAddresses is a slice of string that have been validated as valid network addresses.
// Use it as a flag value when you want to pass a comma-separated list of strings to a flag
// and have it to be automatically parsed into a slice and validated as valid network addresses.
type ListenAddresses []string

// Set implements the flag.Value interface. It parses the string as a comma-separated string.
// Additionally, each value is then passed to net.SplitHostPort for validation that it's a correct network address.
func (a *ListenAddresses) Set(s string) error {
	for _, t := range strings.Split(s, ",") {
		_, _, err := net.SplitHostPort(t)
		if err != nil {
			return err
		}

		*a = append(*a, t)
	}

	return nil
}

// String implements the flag.Value interface. It returns the slice as a comma-separated string.
func (a ListenAddresses) String() string {
	return strings.Join(a, ",")
}

// Strings is a slice of string.
// Use it as a flag value when you want to pass a comma-separated list of strings to a flag
// and have it to be automatically parsed into a slice.
type Strings []string

// Set implements the flag.Value interface. It parses the string as a comma-separated string.
func (s *Strings) Set(str string) error {
	for _, t := range strings.Split(str, ",") {
		*s = append(*s, t)
	}

	return nil
}

// String implements the flag.Value interface. It returns the slice as a comma-separated string.
func (s Strings) String() string {
	return strings.Join(s, ",")
}
