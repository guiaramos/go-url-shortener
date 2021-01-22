package shortener

import (
	"errors"
	"time"

	errs "github.com/pkg/errors"
	"github.com/teris-io/shortid"
	"gopkg.in/dealancer/validate.v2"
)

var (
	ErrRedirectNotFound = errors.New("Redirect Not Found")
	ErrRedirectInvalid  = errors.New("Redirect Invalid")
)

type redirectService struct {
	repo RedirectRepository
}

// NewRedirectService creates a new redirect service
func NewRedirectService(repo RedirectRepository) RedirectService {
	return &redirectService{
		repo,
	}
}

func (s *redirectService) Find(code string) (*Redirect, error) {
	return s.repo.Find(code)
}

func (s *redirectService) Store(r *Redirect) error {
	if err := validate.Validate(r); err != nil {
		return errs.Wrap(ErrRedirectInvalid, "service.Redirect.Store")
	}
	r.Code = shortid.MustGenerate()
	r.CreatedAt = time.Now().UTC().Unix()
	return s.repo.Store(r)
}
