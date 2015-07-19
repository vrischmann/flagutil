package flagutil

import (
	"net"
	"strings"
)

type ListenAddresses []string

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

func (a ListenAddresses) String() string {
	return strings.Join(a, ",")
}

type Strings []string

func (s *Strings) Set(str string) error {
	for _, t := range strings.Split(str, ",") {
		*s = append(*s, t)
	}

	return nil
}

func (s Strings) String() string {
	return strings.Join(s, ",")
}
