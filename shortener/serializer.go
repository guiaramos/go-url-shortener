package shortener

// RedirectSerializer is a interface for implement the Redirect serialization
type RedirectSerializer interface {
	Decode(input []byte) (*Redirect, error)
	Encode(input *Redirect) ([]byte, error)
}
