package extract

import (
  "net/http"
  "errors"
  "io"
  "strings"
  "encoding/json"

  "github.com/PuerkitoBio/goquery"
)

func FromURL(url string) (result []map[string]interface{}, err error) {
	res, err := http.Get(url)
	if err != nil {
		return result, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return result, errors.New(res.Status)
	}
	return FromReader(res.Body)
}

func FromHTML(html string) (result []map[string]interface{}, err error) {
	return FromReader(strings.NewReader(html))
}

func FromReader(r io.Reader) (result []map[string]interface{}, err error) {
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