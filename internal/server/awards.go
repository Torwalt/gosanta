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
	Reason    string    `json:"reason"`
	EmailRef  string    `json:"email_ref"`
}

type awardsHandler struct {
	s ports.AwardReadingService
}

func (h *awardsHandler) router() chi.Router {
	r := chi.NewRouter()

	r.Get("/user", h.getUserAwards)

	return r
}

func (h *awardsHandler) getUserAwards(w http.ResponseWriter, r *http.Request) {
	uIds := chi.URLParam(r, "userID")
	uId, err := strconv.Atoi(uIds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encodeError(err, w)
		return
	}
	awards, err := h.s.GetUserAwards(awards.UserId(uId))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encodeError(err, w)
		return
	}
	resp := []UserAwardResponse{}
	for _, a := range awards {
		uar := UserAwardResponse{
			Id: a.Id, UserId: int(a.AssignedTo), CreatedOn: a.EarnedOn,
			Reason: a.Reason.String(), EmailRef: a.EmailRef}
		resp = append(resp, uar)
	}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encodeError(err, w)
		return
	}
}
