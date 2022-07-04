package time

import (
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

func (d Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

func (d *Duration) UnmarshalJSON(b []byte) error {
	var unmarshalledJSON interface{}

	if err := json.Unmarshal(b, &unmarshalledJSON); err != nil {
		return err
	}

	switch t := unmarshalledJSON.(type) {
	case int, int32, int64, float64:
		d.Duration = time.Duration(t.(float64))
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
