package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/daetal-us/getld/extract"
)

func main() {
	flag.Parse()
	results := []map[string]interface{}{}
	urls := flag.Args()
	if len(urls) == 0 {
		fmt.Println("At least one URL is required.")
		os.Exit(1)
	}
	for _, url := range urls {
		result, err := extract.FromURL(url)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		results = append(results, result...)
	}
	encoded, err := json.Marshal(results)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(encoded))
}
