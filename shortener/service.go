package shortener

// RedirectService is interface for implement the Redirect service
type RedirectService interface {
	Find(code string) (*Redirect, error)
	Store(r *Redirect) error
}
