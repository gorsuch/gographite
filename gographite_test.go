package gographite

import (
	"fmt"
	"log"
)

func ExampleRequestURL() {
	c, err := NewClient("http://graphite.example.com:8000")
	if err != nil {
		log.Fatal(err)
	}

	t := []string{"system.a.load.1m", "foo.bar.baz"}

	fmt.Println(c.RequestURL(t, "-1h"))
	// Output:
	// http://graphite.example.com:8000/render?format=json&from=-1h&target=system.a.load.1m&target=foo.bar.baz
}
