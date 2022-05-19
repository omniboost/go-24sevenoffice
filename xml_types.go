package twentyfour

import (
	"encoding/json"
	"encoding/xml"
	"time"

	"github.com/pkg/errors"
)

type Date struct {
	time.Time
}

func (d Date) MarshalSchema() string {
	return d.Time.Format("2006-01-02")
}

func (d Date) IsEmpty() bool {
	return d.Time.IsZero()
}

func (d Date) MarshalXML(enc *xml.Encoder, start xml.StartElement) error {
	if d.Time.IsZero() {
		return nil
	}
	return enc.EncodeElement(d.Time.Format("2006-01-02T15:04:05"), start)
}

func (d *Date) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	var value string
	err := dec.DecodeElement(&value, &start)
	if err != nil {
		return err
	}

	if value == "" {
		return nil
	}

	// first try standard date
	layout := time.RFC3339
	d.Time, err = time.Parse(layout, value)
	if err == nil {
		return nil
	}

	layout = "2006-01-02"
	d.Time, err = time.Parse(layout, value)
	if err == nil {
		return errors.WithStack(err)
	}

	return nil
}

type DateTime struct {
	time.Time
}

func (dt DateTime) IsEmpty() bool {
	return dt.Time.IsZero()
}

func (dt DateTime) MarshalXML(enc *xml.Encoder, start xml.StartElement) error {
	if dt.Time.IsZero() {
		return nil
	}
	return enc.EncodeElement(dt.Time.Format("2006-01-02T15:04:05"), start)
}

func (d *Date) MarshalJSON() ([]byte, error) {
	if d.Time.IsZero() {
		return json.Marshal(nil)
	}

	return json.Marshal(d.Time.Format("2006-01-02T15:04:05"))
}

func (d *Date) UnmarshalJSON(text []byte) (err error) {
	var value string
	err = json.Unmarshal(text, &value)
	if err != nil {
		return err
	}

	if value == "" {
		return nil
	}

	// first try standard date
	d.Time, err = time.Parse(time.RFC3339, value)
	if err == nil {
		return nil
	}

	d.Time, err = time.Parse("2006-01-02", value)
	if err == nil {
		return nil
	}

	d.Time, err = time.Parse("02/01/2006", value)
	if err == nil {
		return nil
	}

	d.Time, err = time.Parse("02.01.2006", value)
	if err == nil {
		return nil
	}

	d.Time, err = time.Parse("2006-01-02T15:04:05", value)
	return err
}

func (d DateTime) MarshalSchema() string {
	return d.Time.Format(time.RFC3339)
}

func (dt *DateTime) MarshalJSON() ([]byte, error) {
	if dt.Time.IsZero() {
		return json.Marshal(nil)
	}

	return json.Marshal(dt.Time.Format("2006-01-02T15:04:05"))
}

func (dt *DateTime) UnmarshalJSON(text []byte) (err error) {
	var value string
	err = json.Unmarshal(text, &value)
	if err != nil {
		return err
	}

	if value == "" {
		return nil
	}

	// first try standard date
	dt.Time, err = time.Parse(time.RFC3339, value)
	if err == nil {
		return nil
	}

	dt.Time, err = time.Parse("2006-01-02T15:04:05", value)
	return err
}

type Bool bool

func (b *Bool) UnmarshalJSON(text []byte) (err error) {
	var bl bool
	err = json.Unmarshal(text, &bl)
	if err == nil {
		*b = Bool(bl)
		return nil
	}

	var str string
	err = json.Unmarshal(text, &str)
	if err == nil {
		return nil
	}

	if str == "" {
		return nil
	}

	if str == "F" {
		*b = false
	}

	return errors.New("FML")
}
