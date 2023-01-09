package time

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Common durations. There is no definition for units of Day or larger
// to avoid confusion across daylight savings time zone transitions.
//
// To count the number of units in a Duration, divide:
//
//	second := time.Second
//	fmt.Print(int64(second/time.Millisecond)) // prints 1000
//
// To convert an integer number of units to a Duration, multiply:
//
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
	Day                       = 24 * time.Hour
)

var (
	daysRegex = regexp.MustCompile(`^(\d+)d\w*$`)
)

type Duration time.Duration

func (d Duration) Duration() time.Duration {
	return time.Duration(d)
}

func (d Duration) MarshalText() ([]byte, error) {
	return []byte(time.Duration(d).String()), nil
}

func (d *Duration) UnmarshalText(b []byte) error {
	return d.unmarshalText(string(b))
}

func (d Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Duration(d).String())
}

func (d *Duration) UnmarshalJSON(b []byte) error {
	var unmarshalledJSON any
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
		*d = Duration(dur)
	case string:
		return d.unmarshalText(t)
	default:
		return fmt.Errorf("invalid type for duration: %#v", unmarshalledJSON)
	}

	return nil
}

func (d *Duration) unmarshalText(text string) error {
	days := daysRegex.FindAllStringSubmatch(text, -1)
	if len(days) == 0 || len(days[0]) != 2 {
		v, err := time.ParseDuration(text)
		if err != nil {
			return err
		}
		*d = Duration(v)
		return nil
	}

	// Text contains days.
	i, err := strconv.Atoi(days[0][1])
	if err != nil {
		return fmt.Errorf("invalid duration: %w", err)
	}
	*d = Duration(Day * time.Duration(i))

	// Remove days from text and continue the typical way.
	rest := strings.TrimPrefix(text, days[0][1]+"d")
	if len(rest) == 0 {
		return nil
	}

	v, err := time.ParseDuration(rest)
	if err != nil {
		return err
	}
	*d += Duration(v)

	return nil
}
