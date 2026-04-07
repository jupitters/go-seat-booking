package booking

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jupitters/go-seat-booking/internal/utils"
)

type handler struct {
	svc *Service
}

func NewHandler(svc *Service) *handler {
	return &handler{svc}
}

func (h *handler) HoldSeat(w http.ResponseWriter, r *http.Request) {
	type holdRequest struct {
		UserID string `json:"user_id"`
	}

	type holdResponse struct {
		SessionID string `json:"session_id"`
		MovieID   string `json:"movie_id"`
		SeatID    string `json:"movie_id"`
		ExpiresAt string `json:"expires_at"`
	}

	movieID := r.PathValue("movieID")
	seatID := r.PathValue("seatID")

	var req holdRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(err)
		return
	}

	data := Booking{
		UserID:  req.UserID,
		MovieID: movieID,
		SeatID:  seatID,
	}

	err := h.svc.Book(data)
	if err != nil {
		log.Println(err)
		return
	}
}

func (h *handler) ListSeats(w http.ResponseWriter, r *http.Request) {
	movieID := r.PathValue("movieID")

	bookings := h.svc.store.ListBookings(movieID)

	seats := make([]seatInfo, 0, len(bookings))
	for _, b := range bookings {
		seats = append(seats, seatInfo{
			SeatID: b.SeatID,
			UserID: b.UserID,
			Booked: true,
		})
	}

	utils.WriteJson(w, http.StatusOK, seats)
}

type seatInfo struct {
	SeatID string `json:"seat_id"`
	UserID string `json:"user_id"`
	Booked bool   `json:"booked"`
}
