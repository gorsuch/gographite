package gographite

import (
	"fmt"
	"log"
	"net/url"
)

func ExampleRequestURL() {
	url, err := url.Parse("http://graphite.example.com:8000")
	if err != nil {
		log.Fatal(err)
	}

	c := Client{BaseURL: url}
	t := []string{"system.a.load.1m", "foo.bar.baz"}

	fmt.Println(c.RequestURL(t, "-1h"))
	// Output:
	// http://graphite.example.com:8000/render?format=json&from=-1h&target=system.a.load.1m&target=foo.bar.baz
}
