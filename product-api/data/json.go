package data

import (
	"encoding/json"
	"io"
)

// ToJSON serializes the given v (any type) into a string based JSON format
// and then write to an io.Writer
func ToJSON(v any, w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(v)
}

// FromJSON deserializes the object from JSON string
// in an io.Reader to given v (any type).
func FromJSON(v any, r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(v)
}
