# gographite [![Build Status](https://travis-ci.org/gorsuch/gographite.svg?branch=master)](https://travis-ci.org/gorsuch/gographite)

Lightweight go client library for the graphite render api

Fetches multiple targets as JSON, munges to make the results slightly more useful.  Also, `nulls` are replaced with zeros.

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
results, err := c.Render(targets, "-5min")
if err != nil {
  log.Fatal(err)
}

fmt.Println(results)
```

# graphitec

A small command line utility to query graphite metrics

## Installation

```sh
$ go get github.com/gorsuch/gographite/cmd/graphitec
```

## Usage

```sh
$ graphitec -graphite http://graphite.example.com -target carbon.agents.*.metricsReceived
```
