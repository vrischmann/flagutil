package flagutil_test

import (
	"flag"
	"fmt"
	"os"

	"github.com/vrischmann/flagutil"
)

func ExampleNetworkAddress() {
	var addr flagutil.NetworkAddress
	flag.Var(&addr, "h", "Address")

	os.Args = append(os.Args, "-h", "localhost:4000")
	flag.Parse()

	fmt.Println(addr)
	// Output:
	// localhost:4000
}

func ExampleNetworkAddresses() {
	var addrs flagutil.NetworkAddresses
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

func ExampleURL() {
	var url flagutil.URL
	flag.Var(&url, "u", "URL")

	os.Args = append(os.Args, "-u", "https://google.com")
	flag.Parse()

	fmt.Println(url.URL.Scheme)
	fmt.Println(url.URL.Host)
	fmt.Println(url.String())
	// Output:
	// https
	// google.com
	// https://google.com
}

func ExampleURLs() {
	var urls flagutil.URLs
	flag.Var(&urls, "U", "URLs")

	os.Args = append(os.Args, "-U", "https://google.com,https://google.de")
	flag.Parse()

	fmt.Println(urls[0].String())
	fmt.Println(urls[1].String())
	// Output:
	// https://google.com
	// https://google.de
}
