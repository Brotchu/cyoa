package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

func main() {
	filename := flag.String("file", "gopher.json", "The JSON file with CYOA story")
	flag.Parse()
	fmt.Printf("Using story in %s.\n", *filename) //TODO: check flag package

	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}
	d := json.NewDecoder(f)
	var story cyoa.Story
	if err := d.Decode(&story); err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", story)
}
