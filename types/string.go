package types

import (
	"encoding/xml"
	"strings"
)

// FseString represents strings where "null" or similar are interpreted as empty string
type FseString string

// UnmarshalXML parses a string with Null or similar values into a string.
func (t *FseString) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var value string
	err := d.DecodeElement(&value, &start)
	if err != nil {
		return err
	}

	// make sure no surrounding (or remaining) white space pollute the result
	value = strings.TrimSpace(value)

	// an empty field is not an error
	if value == "" {
		*t = FseString("")
		return nil
	}
	check := strings.ToUpper(value)
	switch check {
	case "NIL", "NULL", "VOID":
		*t = FseString("")
	default:
		*t = FseString(value)
	}
	return nil
}

// MarshalXML formats an empty string as "" in order to keep compatibility with
// Go's encoding/xml default behaviour.
func (t *FseString) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if t == nil {
		return e.EncodeElement("", start)
	}
	value := string(*t)
	return e.EncodeElement(value, start)
}
