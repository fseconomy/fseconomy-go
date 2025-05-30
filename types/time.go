package types

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// FseHobbsTime represents Hobbs meter time in "HH:MM" format,
// internally stored as float64 (e.g. "51:30" → 51.5).
type FseHobbsTime float64

// UnmarshalXML parses a Hobbs time string in the format "HH:MM" into a float64.
func (t *FseHobbsTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var value string
	err := d.DecodeElement(&value, &start)
	if err != nil {
		return err
	}

	// make sure no surrounding (or remaining) white space pollute the result
	value = strings.TrimSpace(value)

	// an empty field is not an error
	if value == "" {
		*t = FseHobbsTime(0.0)
		return nil
	}

	parts := strings.Split(value, ":")
	if len(parts) != 2 {
		return fmt.Errorf("invalid hobbs format: %q (expected HH:MM)", value)
	}

	hours, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("invalid hours in hobbs format '%s': %w", value, err)
	}

	minutes, err := strconv.Atoi(parts[1])
	if err != nil {
		return fmt.Errorf("invalid minutes in hobbs format '%s': %w", value, err)
	}

	if minutes < 0 || minutes > 59 {
		return fmt.Errorf("minutes '%d' outside of valid range (0-59) in hobbs time '%s'", minutes, value)
	}
	*t = FseHobbsTime(float64(hours) + float64(minutes)/60.0)
	return nil
}

// MarshalXML formats the Hobbs time as "HH:MM".
func (t *FseHobbsTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if t == nil {
		return e.EncodeElement("00:00", start)
	}

	value := float64(*t)
	hoursPart := int(value) // full hours
	minutesFraction := (value - float64(hoursPart)) * 60.0

	// round minutes to avoid display issues
	minutesPart := int(minutesFraction + 0.5)
	if minutesPart >= 60 { // carry forward if minutes are rounded up to 60
		hoursPart++
		minutesPart = 0
	}
	return e.EncodeElement(fmt.Sprintf("%02d:%02d", hoursPart, minutesPart), start)
}

// FseAssignmentTimeLayout defines the layout used for assignment time parsing.
const FseAssignmentTimeLayout = "2006-01-02 15:04:05"

// FseAssignmentTime represents a timestamp in "2006-01-02 15:04:05" format.
type FseAssignmentTime struct {
	time.Time
}

// UnmarshalXML parses the assignment time from XML using a fixed layout.
func (t *FseAssignmentTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	parsedTime, err := unmarshalXMLTime(d, start, FseAssignmentTimeLayout)
	if err != nil {
		return err
	}
	t.Time = parsedTime
	return nil
}

// MarshalXML formats the assignment time using the fixed layout.
func (t *FseAssignmentTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if t.Time.IsZero() {
		return e.EncodeElement("", start)
	}
	return e.EncodeElement(t.Time.Format(FseAssignmentTimeLayout), start)
}

// FseFlightLogTimeLayout defines the layout used for flight log time parsing.
const FseFlightLogTimeLayout = "2006/01/02 15:04:05"

// FseFlightLogTime represents a timestamp in "2006/01/02 15:04:05" format.
type FseFlightLogTime struct {
	time.Time
}

// UnmarshalXML parses the flight log time from XML using a fixed layout.
func (t *FseFlightLogTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	parsedTime, err := unmarshalXMLTime(d, start, FseFlightLogTimeLayout)
	if err != nil {
		return err
	}
	t.Time = parsedTime
	return nil
}

// MarshalXML formats the flight log time using the fixed layout.
func (t *FseFlightLogTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if t.Time.IsZero() {
		return e.EncodeElement("", start)
	}
	return e.EncodeElement(t.Time.Format(FseFlightLogTimeLayout), start)
}

// unmarshalXMLTime is a helper to parse time values with custom layouts from XML.
func unmarshalXMLTime(d *xml.Decoder, start xml.StartElement, layout string) (time.Time, error) {
	var value string
	err := d.DecodeElement(&value, &start)
	if err != nil {
		return time.Time{}, err
	}

	// make sure no surrounding (or remaining) white space pollute the result
	value = strings.TrimSpace(value)

	// an empty field is not an error
	if value == "" {
		return time.Time{}, nil
	}

	t, err := time.Parse(layout, value)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}
