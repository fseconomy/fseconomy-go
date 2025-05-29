package types

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestFseCommodityAmount_UnmarshalXML(t *testing.T) {
	xmlData := `<Root><Amount>123 kg</Amount></Root>`
	var parsed struct {
		Amount FseCommodityAmount `xml:"Amount"`
	}
	err := xml.Unmarshal([]byte(xmlData), &parsed)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}
	expected := 123
	if int(parsed.Amount) != expected {
		t.Errorf("Expected %d, got %d", expected, parsed.Amount)
	}
}

func TestFseCommodityAmount_MarshalXML(t *testing.T) {
	type amountWrapper struct {
		XMLName xml.Name           `xml:"Root"`
		Amount  FseCommodityAmount `xml:"Amount"`
	}
	data := amountWrapper{
		Amount: 123,
	}
	output, err := xml.Marshal(&data)
	if err != nil {
		t.Fatalf("Failed to marshal: %v", err)
	}
	expected := `<Root><Amount>123 kg</Amount></Root>`
	if !strings.Contains(string(output), expected) {
		t.Errorf("Expected %q in output, got %q", expected, output)
	}
}
