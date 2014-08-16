gographite
==========

Lightweight go client library for the graphite render api

## Installation

```sh
$ go get github.com/gorsuch/gographite
```

## Usage

```go
c, err := NewClient("http://graphite.example.com:8000")
if err != nil {
	log.Fatal(err)
}

targets := []string{"system.a.load.1m", "foo.bar.baz"}
results, err := c.Render(targets)
if err != nil {
  log.Fatal(err)
}

fmt.Println(results)
```
