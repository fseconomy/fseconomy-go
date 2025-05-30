package types

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestFseString_UnmarshalXML(t *testing.T) {
	xmlData := `<Root><Text>Null</Text></Root>`
	var parsed struct {
		Amount FseString `xml:"Text"`
	}
	err := xml.Unmarshal([]byte(xmlData), &parsed)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}
	expected := ""
	if string(parsed.Amount) != expected {
		t.Errorf("Expected %s, got %s", expected, parsed.Amount)
	}
}

func TestFseString_MarshalXML(t *testing.T) {
	type textWrapper struct {
		XMLName xml.Name  `xml:"Root"`
		Text    FseString `xml:"Text"`
	}
	data := textWrapper{
		Text: "",
	}
	output, err := xml.Marshal(&data)
	if err != nil {
		t.Fatalf("Failed to marshal: %v", err)
	}
	expected := `<Root><Text></Text></Root>`
	if !strings.Contains(string(output), expected) {
		t.Errorf("Expected %q in output, got %q", expected, output)
	}
}
