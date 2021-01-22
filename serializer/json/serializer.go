package json

import (
	"encoding/json"

	"github.com/guiaramos/go-url-shortener/shortener"
	"github.com/pkg/errors"
)

// Redirect is a collection of methods to implement the redirect serialize JSON
type Redirect struct{}

// Decode unmarshal the input
func (r Redirect) Decode(input []byte) (*shortener.Redirect, error) {
	redirect := &shortener.Redirect{}
	if err := json.Unmarshal(input, redirect); err != nil {
		return nil, errors.Wrap(err, "serializer.Redirect.Decode")
	}
	return redirect, nil
}

// Encode marshal the input
func (r Redirect) Encode(input *shortener.Redirect) ([]byte, error) {
	msg, err := json.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.Redirect.Encode")
	}
	return msg, nil
}
