package shortener_test

import (
	"testing"

	"github.com/guiaramos/go-url-shortener/shortener"
)

type SpyRedirectRepository struct {
	redirect *shortener.Redirect
}

func (s SpyRedirectRepository) Find(code string) (*shortener.Redirect, error) {
	return s.redirect, nil
}

func (s *SpyRedirectRepository) Store(r *shortener.Redirect) error {
	s.redirect = r
	return nil
}

func TestRedirectService(t *testing.T) {
	repo := &SpyRedirectRepository{}
	service := shortener.NewRedirectService(repo)

	t.Run("should create a new redirect", func(t *testing.T) {
		r := &shortener.Redirect{URL: "http://example.com/"}
		err := service.Store(r)
		assertNoError(t, err)

		nr, err := service.Find("fakecode")
		assertNoError(t, err)

		if nr.Code == "" {
			t.Fatal("expected to have a Code but didn't get one")
		}
	})
}

func assertNoError(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("expected no error but got %v\n", err)
	}
}
