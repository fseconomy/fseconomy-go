package types

import (
	"encoding/xml"
	"fmt"
	"strings"
)

// FseBool represents a boolean flag provided as "True" or "False", but
// also understanding other "dialects" (such as "Yes", "No", "On", "Off", "Aye", "Nay")
// internally stored as bool (e.g. "Yes" → true).
type FseBool bool

// UnmarshalXML parses a yes/no flag in the format "True" or "False" into a bool.
func (t *FseBool) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var value string
	err := d.DecodeElement(&value, &start)
	if err != nil {
		return err
	}

	// make sure no surrounding (or remaining) white space pollute the result
	value = strings.TrimSpace(value)

	// an empty field is not an error
	if value == "" {
		*t = FseBool(false)
		return nil
	}
	value = strings.ToUpper(value)
	err = nil
	switch value {
	case "YES", "yes", "TRUE", "true", "ON", "on", "AYE", "aye":
		*t = FseBool(true)
	case "NO", "no", "FALSE", "false", "OFF", "off", "NAY", "nay":
		*t = FseBool(false)
	default:
		*t = FseBool(false)
		err = fmt.Errorf("unexpected flag '%s' - don't know what to do with this one", value)
	}
	return err
}

// MarshalXML formats the flag as "True" or "False" in order to keep compatibility with
// Go's encoding/xml default behaviour.
func (t *FseBool) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if t == nil {
		return e.EncodeElement("False", start)
	}
	value := bool(*t)
	if value {
		return e.EncodeElement("True", start)
	}
	return e.EncodeElement("False", start)
}
