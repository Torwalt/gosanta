package rest

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	awards "gosanta/internal"
	"gosanta/internal/ports"

	"github.com/go-chi/chi"
)

type UserAwardResponse struct {
	Id        int64     `json:"id"`
	UserId    int       `json:"user_id"`
	CreatedOn time.Time `json:"created_on"`
	Type      string    `json:"award_type"`
}

type LeaderboardMember struct {
	UserId       int    `json:"user_id"`
	UserFullName string `json:"user_full_name"`
	Score        int    `json:"score"`
	IgnoreCount  int    `json:"ignore_count"`
	OpenCount    int    `json:"open_count"`
	ReportCount  int    `json:"report_count"`
	Rank         int    `json:"rank"`
}

type awardsHandler struct {
	s ports.AwardReadingService
}

func (h *awardsHandler) router() chi.Router {
	r := chi.NewRouter()

	r.Route("/user", func(r chi.Router) {
		r.Route("/{userID}", func(r chi.Router) {
			r.Use()
			r.Get("/", h.getUserAwards)
			r.Route("/leaderboard", func(r chi.Router) {
				r.Get("/", h.calcLeaderboard)
			})
		})
	})

	return r
}

func (h *awardsHandler) getUserAwards(w http.ResponseWriter, r *http.Request) {
	uIdStr := chi.URLParam(r, "userID")
	uId, err := strconv.Atoi(uIdStr)
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
			Type: a.Type.String(),
		}
		resp = append(resp, uar)
	}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		encodeError(err, w)
		return
	}
}

func (h *awardsHandler) calcLeaderboard(w http.ResponseWriter, r *http.Request) {
	uIds := chi.URLParam(r, "userID")
	uId, err := strconv.Atoi(uIds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeError(err, w)
		return
	}
	lb, err := h.s.CalcLeaderboard(awards.UserId(uId))
	if err != nil {
		encodeError(err, w)
		return
	}

	lbr := []LeaderboardMember{}
	for _, member := range lb.RankedUsers {
		lm := LeaderboardMember{
			UserId:       int(member.UserId),
			UserFullName: member.UserFullName,
			Score:        member.Score,
			IgnoreCount:  member.Summary.IgnoringAward,
			OpenCount:    member.Summary.OpenAward,
			ReportCount:  member.Summary.ReportAward,
			Rank:         member.Rank,
		}
		lbr = append(lbr, lm)
	}
	if err := json.NewEncoder(w).Encode(lbr); err != nil {
		encodeError(err, w)
		return
	}
}
