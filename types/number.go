package types

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

// FseCommodityAmount represents commodity amounts provided as "x kg",
// internally stored as int (e.g. "512 kg" → 512).
type FseCommodityAmount int

// UnmarshalXML parses a commodity amount string in the format "n kg" into an int.
func (t *FseCommodityAmount) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var value string
	err := d.DecodeElement(&value, &start)
	if err != nil {
		return err
	}

	// make sure no surrounding (or remaining) white space pollute the result
	value = strings.TrimSpace(value)

	// an empty field is not an error
	if value == "" {
		*t = FseCommodityAmount(0)
		return nil
	}

	parts := strings.Split(value, " ")
	if len(parts) != 2 {
		return fmt.Errorf("invalid amount format: %q (expected n kg)", value)
	}

	amount, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("invalid amount in commodity amount format '%s': %w", value, err)
	}

	*t = FseCommodityAmount(amount)
	return nil
}

// MarshalXML formats the commodity amount as "n kg".
func (t *FseCommodityAmount) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if t == nil {
		return e.EncodeElement("0 kg", start)
	}
	value := int(*t)
	return e.EncodeElement(fmt.Sprintf("%d kg", value), start)
}
