package types

import (
	"encoding/xml"
	"log"
	"strings"
	"testing"
	"time"
)

func TestFseHobbsTime_UnmarshalXML(t *testing.T) {
	xmlData := `<Root><AirframeTime>12:30</AirframeTime></Root>`
	var parsed struct {
		AirframeTime FseHobbsTime `xml:"AirframeTime"`
	}
	err := xml.Unmarshal([]byte(xmlData), &parsed)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}
	log.Printf("FseHobbsTime UnmarshalXML: %v", parsed)
	expected := 12.5
	if float64(parsed.AirframeTime) != expected {
		t.Errorf("Expected %.2f, got %.2f", expected, parsed.AirframeTime)
	}
}

func TestFseHobbsTime_MarshalXML(t *testing.T) {
	type hobbsWrapper struct {
		XMLName      xml.Name     `xml:"Root"`
		AirframeTime FseHobbsTime `xml:"AirframeTime"`
	}
	data := hobbsWrapper{
		AirframeTime: 2.75,
	}
	output, err := xml.Marshal(&data)
	if err != nil {
		t.Fatalf("Failed to marshal: %v", err)
	}
	expected := `<Root><AirframeTime>02:45</AirframeTime></Root>`
	if !strings.Contains(string(output), expected) {
		t.Errorf("Expected %q in output, got %q", expected, output)
	}
}

func TestFseAssignmentTime_UnmarshalXML(t *testing.T) {
	xmlData := `<Root><Time>2025-06-02 14:30:22</Time></Root>`
	var parsed struct {
		Time FseAssignmentTime `xml:"Time"`
	}
	err := xml.Unmarshal([]byte(xmlData), &parsed)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}
	expected := time.Date(2025, 6, 2, 14, 30, 22, 0, time.UTC)
	if !parsed.Time.Equal(expected) {
		t.Errorf("Expected %v, got %v", expected, parsed.Time)
	}
}

func TestFseFlightLogTime_UnmarshalXML(t *testing.T) {
	xmlData := `<Root><Time>2025/06/02 14:30:22</Time></Root>`
	var parsed struct {
		Time FseFlightLogTime `xml:"Time"`
	}
	err := xml.Unmarshal([]byte(xmlData), &parsed)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}
	expected := time.Date(2025, 6, 2, 14, 30, 22, 0, time.UTC)
	if !parsed.Time.Equal(expected) {
		t.Errorf("Expected %v, got %v", expected, parsed.Time)
	}
}
