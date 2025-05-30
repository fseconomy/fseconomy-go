package types

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestFseBool_UnmarshalXML(t *testing.T) {
	xmlData := `<Root><Switch>Yes</Switch></Root>`
	var parsed struct {
		Switch FseBool `xml:"Switch"`
	}
	err := xml.Unmarshal([]byte(xmlData), &parsed)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}
	if !parsed.Switch {
		t.Errorf("Expected %t, got %t", true, parsed.Switch)
	}
}

func TestFseBool_MarshalXML(t *testing.T) {
	type switchWrapper struct {
		XMLName xml.Name `xml:"Root"`
		Switch  FseBool  `xml:"Switch"`
	}
	data := switchWrapper{
		Switch: true,
	}
	output, err := xml.Marshal(&data)
	if err != nil {
		t.Fatalf("Failed to marshal: %v", err)
	}
	expected := `<Root><Switch>True</Switch></Root>`
	if !strings.Contains(string(output), expected) {
		t.Errorf("Expected %q in output, got %q", expected, output)
	}
}
