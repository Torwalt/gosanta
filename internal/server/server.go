package server

import (
	"encoding/json"
	"gosanta/internal/ports"
	"net/http"

	"github.com/go-chi/chi"
)

type Server struct {
	awardService ports.AwardReadingService

	router chi.Router
}

func New(awardService ports.AwardReadingService) Server {
	return Server{awardService: awardService}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func encodeError(err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	// switch err {
	// case shipping.ErrUnknownCargo:
	// 	w.WriteHeader(http.StatusNotFound)
	// case tracking.ErrInvalidArgument:
	// 	w.WriteHeader(http.StatusBadRequest)
	// default:
	// 	w.WriteHeader(http.StatusInternalServerError)
	// }
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
