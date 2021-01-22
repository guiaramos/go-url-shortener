package msgpack

import (
	"github.com/guiaramos/go-url-shortener/shortener"
	"github.com/pkg/errors"
	"github.com/vmihailenco/msgpack"
)

// Redirect is a collection of methods to implement serialize MsgPack
type Redirect struct{}

// Decode unmarshal MsgPack
func (r Redirect) Decode(input []byte) (*shortener.Redirect, error) {
	redirect := &shortener.Redirect{}
	if err := msgpack.Unmarshal(input, redirect); err != nil {
		return nil, errors.Wrap(err, "serializer.Redirect.Decode")
	}
	return redirect, nil
}

// Encode marshal MsgPack
func (r Redirect) Encode(input *shortener.Redirect) ([]byte, error) {
	msg, err := msgpack.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serialize.Redirect.Encode")
	}
	return msg, nil
}
