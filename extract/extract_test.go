package extract

import (
	"testing"
  "encoding/json"
)

func TestExtractFromHTML(t *testing.T) {
	expected := `{"@type":"organization","name":"example"}`
	html := `<html><body><script type="application/ld+json">` + expected + `</script></body></html>"`

	extracted, err := FromHTML(html)
	if err != nil {
		t.Fatal(err)
	}

	encoded, err := json.Marshal(extracted[0])
	if err != nil {
		t.Fatal(err)
	}

	result := string(encoded)
	if result != expected {
		t.Errorf("Unxpected encoded result: %s", result)
	}
}
