package time

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// Common durations. There is no definition for units of Day or larger
// to avoid confusion across daylight savings time zone transitions.
//
// To count the number of units in a Duration, divide:
//	second := time.Second
//	fmt.Print(int64(second/time.Millisecond)) // prints 1000
//
// To convert an integer number of units to a Duration, multiply:
//	seconds := 10
//	fmt.Print(time.Duration(seconds)*time.Second) // prints 10s
//
// NOTE: This is a legacy compatibility layer for the time package.
const (
	Nanosecond  time.Duration = 1
	Microsecond               = 1000 * time.Nanosecond
	Millisecond               = 1000 * time.Microsecond
	Second                    = 1000 * time.Millisecond
	Minute                    = 60 * time.Second
	Hour                      = 60 * time.Minute
)

type Duration struct {
	time.Duration
}

func (d Duration) MarshalText() ([]byte, error) {
	return []byte(d.String()), nil
}

func (d *Duration) UnmarshalText(b []byte) error {
	v, err := time.ParseDuration(string(b))
	if err != nil {
		return err
	}
	d.Duration = v
	return nil
}

func (d Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

func (d *Duration) UnmarshalJSON(b []byte) error {
	var unmarshalledJSON interface{}
	decoder := json.NewDecoder(bytes.NewReader(b))
	decoder.UseNumber()
	if err := decoder.Decode(&unmarshalledJSON); err != nil {
		return err
	}

	switch t := unmarshalledJSON.(type) {
	case json.Number:
		dur, err := t.Int64()
		if err != nil {
			return err
		}
		d.Duration = time.Duration(dur)
	case string:
		v, err := time.ParseDuration(t)
		if err != nil {
			return err
		}
		d.Duration = v
	default:
		return fmt.Errorf("invalid type for duration: %#v", unmarshalledJSON)
	}

	return nil
}
