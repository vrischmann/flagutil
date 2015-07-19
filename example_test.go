package flagutil_test

import (
	"flag"
	"fmt"
	"os"

	"github.com/vrischmann/flagutil"
)

func ExampleListenAddresses() {
	var addrs flagutil.ListenAddresses
	flag.Var(&addrs, "H", "Addresses")

	os.Args = append(os.Args, "-H", "localhost:4000,localhost:5000")
	flag.Parse()

	fmt.Println(addrs[0])
	fmt.Println(addrs[1])
	fmt.Println(addrs)
	// Output:
	// localhost:4000
	// localhost:5000
	// localhost:4000,localhost:5000
}

func ExampleStrings() {
	var strings flagutil.Strings
	flag.Var(&strings, "s", "Strings")

	os.Args = append(os.Args, "-s", "foo,bar")
	flag.Parse()

	fmt.Println(strings[0])
	fmt.Println(strings[1])
	fmt.Println(strings)
	// Output:
	// foo
	// bar
	// foo,bar
}
