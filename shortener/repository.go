package shortener

// RedirectRepository is an interface for implement the Redirect repository
type RedirectRepository interface {
	Find(code string) (*Redirect, error)
	Store(r *Redirect) error
}
