package shotener

// RedirectService is a collection of methods for the Redirect Service
type RedirectService interface {
	Find(code string) (*Redirect, error)
	Store(redirect *Redirect) error
}
