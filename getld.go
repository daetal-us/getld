package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	flag.Parse()
	results := []map[string]interface{}{}
	for _, url := range flag.Args() {
		result, err := ExtractFromURL(url)
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

func ExtractFromURL(url string) (result []map[string]interface{}, err error) {
	res, err := http.Get(url)
	if err != nil {
		return result, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return result, errors.New(res.Status)
	}
	return ExtractFromReader(res.Body)
}

func ExtractFromHTML(html string) (result []map[string]interface{}, err error) {
	return ExtractFromReader(strings.NewReader(html))
}

func ExtractFromReader(r io.Reader) (result []map[string]interface{}, err error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return result, err
	}
	selector := `script[type="application/ld+json"]`
	doc.Find(selector).Each(func(i int, s *goquery.Selection) {
		var decoded map[string]interface{}
		err = json.Unmarshal([]byte(s.Text()), &decoded)
		if err == nil {
			result = append(result, decoded)
		}
	})
	return result, nil
}
