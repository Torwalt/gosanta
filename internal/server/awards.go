package server

import (
	"encoding/json"
	awards "gosanta/internal"
	"gosanta/internal/ports"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
)

type UserAwardResponse struct {
	Id        int64     `json:"id"`
	UserId    int       `json:"user_id"`
	CreatedOn time.Time `json:"created_on"`
	Type      string    `json:"award_type"`
}

type awardsHandler struct {
	s ports.AwardReadingService
}

func (h *awardsHandler) router() chi.Router {
	r := chi.NewRouter()

	r.Route("/user", func(r chi.Router) {
		r.Route("/{userID}", func(r chi.Router) {
			r.Get("/", h.getUserAwards)
		})
	})

	return r
}

func (h *awardsHandler) getUserAwards(w http.ResponseWriter, r *http.Request) {
	uIds := chi.URLParam(r, "userID")
	uId, err := strconv.Atoi(uIds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeError(err, w)
		return
	}
	awardS, err := h.s.GetUserAwards(awards.UserId(uId))
	if err != nil {
		encodeError(err, w)
		return
	}
	resp := []UserAwardResponse{}
	for _, a := range awardS {
		uar := UserAwardResponse{
			Id: a.Id, UserId: int(a.AssignedTo), CreatedOn: a.EarnedOn,
			Type: a.Type.String()}
		resp = append(resp, uar)
	}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		encodeError(err, w)
		return
	}
}
