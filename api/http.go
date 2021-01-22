package api

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	js "github.com/guiaramos/go-url-shortener/serializer/json"
	ms "github.com/guiaramos/go-url-shortener/serializer/msgpack"
	"github.com/guiaramos/go-url-shortener/shortener"
	"github.com/pkg/errors"
)

// RedirectHandler is a collection of methods for REST API
type RedirectHandler interface {
	Get(http.ResponseWriter, *http.Request)
	Post(http.ResponseWriter, *http.Request)
}

type handler struct {
	service shortener.RedirectService
}

// NewHandler creates a new Redirect Handler
func NewHandler(s shortener.RedirectService) RedirectHandler {
	return &handler{service: s}
}

func setupResponse(w http.ResponseWriter, contentType string, body []byte, statusCode int) {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(statusCode)
	_, err := w.Write(body)
	if err != nil {
		log.Println(err)
	}
}

func (h *handler) serializer(contentType string) shortener.RedirectSerializer {
	if contentType == "application/x-msgpack" {
		return &ms.Redirect{}
	}
	return &js.Redirect{}
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")

	redirect, err := h.service.Find(code)
	if err != nil {
		if errors.Cause(err) == shortener.ErrRedirectNotFound {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusNotFound)
			return
		}
		internalServerError(w)
		return
	}

	http.Redirect(w, r, redirect.URL, http.StatusMovedPermanently)
}

func (h *handler) Post(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		internalServerError(w)
		return
	}

	redirect, err := h.serializer(contentType).Decode(requestBody)
	if err != nil {
		internalServerError(w)
		return
	}

	err = h.service.Store(redirect)
	if err != nil {
		if errors.Cause(err) == shortener.ErrRedirectNotFound {
			internalServerError(w)
			return
		}
		internalServerError(w)
		return
	}

	responseBody, err := h.serializer(contentType).Encode(redirect)
	if err != nil {
		internalServerError(w)
		return
	}

	setupResponse(w, contentType, responseBody, http.StatusCreated)
}

func internalServerError(w http.ResponseWriter) {
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
