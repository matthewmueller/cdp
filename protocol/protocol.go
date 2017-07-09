// Code generated by cdpgen. DO NOT EDIT.

package protocol

import (
	"encoding/json"
	"errors"
	"time"
)

// Timestamp represents a timestamp (since epoch).
type Timestamp float64

// String calls (time.Time).String().
func (t Timestamp) String() string {
	return t.Time().String()
}

// Time parses the Unix time with millisecond accuracy.
func (t Timestamp) Time() time.Time {
	secs := int64(t)
	// The Unix time in t only has ms accuracy.
	ms := int64((float64(t) - float64(secs)) * 1000000)
	return time.Unix(secs, ms*1000)
}

// MarshalJSON implements json.Marshaler. Encodes to null if t is zero.
func (t Timestamp) MarshalJSON() ([]byte, error) {
	if t == 0 {
		return []byte("null"), nil
	}
	f := float64(t)
	return json.Marshal(&f)
}

// UnmarshalJSON implements json.Unmarshaler.
func (t *Timestamp) UnmarshalJSON(data []byte) error {
	*t = 0
	if len(data) == 0 {
		return nil
	}
	var f float64
	if err := json.Unmarshal(data, &f); err != nil {
		return errors.New("protocol.Timestamp: " + err.Error())
	}
	*t = Timestamp(f)
	return nil
}

var _ json.Marshaler = (*Timestamp)(nil)
var _ json.Unmarshaler = (*Timestamp)(nil)