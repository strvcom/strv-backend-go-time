package time

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDuration_MarshalText(t *testing.T) {
	d := Duration{Duration: time.Hour * 3}
	expected := []byte("3h0m0s")

	data, err := d.MarshalText()
	assert.NoError(t, err)
	assert.Equal(t, expected, data)
}

func TestDuration_UnmarshalText(t *testing.T) {
	text := "3h0m0s"
	d := Duration{}

	err := d.UnmarshalText([]byte(text))
	assert.NoError(t, err)
	assert.Equal(t, text, d.String())

	err = d.UnmarshalText([]byte("unknown"))
	assert.EqualError(t, err, `time: invalid duration "unknown"`)
}

func TestDuration_MarshalJSON(t *testing.T) {
	d := Duration{Duration: time.Hour * 3}
	expected := []byte(`"3h0m0s"`)

	data, err := d.MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, expected, data)
}

func TestDuration_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name        string
		text        string
		expected    time.Duration
		expectedErr string
	}{
		{
			name:        "success:string",
			text:        `"3h0m0s"`,
			expected:    time.Hour * 3,
			expectedErr: "",
		},
		{
			name:        "success:number",
			text:        "123456789",
			expected:    time.Duration(123456789),
			expectedErr: "",
		},
		{
			name:        "failure:invalid-number",
			text:        "123456789.1",
			expected:    time.Duration(0),
			expectedErr: `strconv.ParseInt: parsing "123456789.1": invalid syntax`,
		},
		{
			name:        "failure:parse-duration",
			text:        `"abc123"`,
			expected:    time.Duration(0),
			expectedErr: `time: invalid duration "abc123"`,
		},
		{
			name:        "failure:json-unmarshal",
			text:        `"{`,
			expected:    time.Duration(0),
			expectedErr: "unexpected EOF",
		},
		{
			name:        "failure:unknown-type",
			text:        `{"array":[]}`,
			expected:    time.Duration(0),
			expectedErr: `invalid type for duration: map[string]interface {}{"array":[]interface {}{}}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Duration{}
			err := d.UnmarshalJSON([]byte(tt.text))
			if tt.expectedErr != "" {
				assert.Equal(t, tt.expectedErr, err.Error())
			}
			assert.Equal(t, tt.expected, d.Duration)
		})
	}
}
