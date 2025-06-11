package rest

import (
	"encoding/json"
	"errors"
	"net/http"

	awards "gosanta/internal"
	"gosanta/internal/ports"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Server struct {
	awardService ports.AwardReadingService
	router       chi.Router
}

func New(awardService ports.AwardReadingService) Server {
	s := Server{awardService: awardService}

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/awards", func(r chi.Router) {
		h := awardsHandler{s: s.awardService}
		r.Mount("/v1", h.router())
	})

	s.router = r

	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func encodeError(err error, w http.ResponseWriter) {
	var awardErr *awards.Error

	if errors.As(err, &awardErr) != true {
		w.WriteHeader(http.StatusInternalServerError)
		writeError(err, w)
		return
	}

	switch awardErr.Code {
	case awards.DoesNotExistError:
		w.WriteHeader(http.StatusNotFound)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}

	writeError(awardErr, w)
}

func writeError(err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
