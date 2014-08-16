package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/gorsuch/gographite"
)

var (
	graphite string
	target   string
	from     string
)

func init() {
	flag.StringVar(&graphite, "graphite", "http://localhost:8080", "location of your graphite api")
	flag.StringVar(&target, "target", "carbon.agents.*.metricsReceived", "target to query")
	flag.StringVar(&from, "from", "-1h", "timespan")
}

func main() {
	flag.Parse()

	c, err := gographite.NewClient(graphite)
	if err != nil {
		log.Fatal(err)
	}

	results, err := c.Render([]string{target}, from)
	if err != nil {
		log.Fatal(err)
	}

	json, err := json.Marshal(results)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(json))
}
