package booking

import (
	"net/http"

	"github.com/jupitters/go-seat-booking/internal/utils"
)

type handler struct {
	svc *Service
}

func NewHandler(svc *Service) *handler {
	return &handler{svc}
}

func (h *handler) ListSeats(w http.ResponseWriter, r *http.Request) {
	movieID := r.PathValue("movieID")

	h.svc.store.ListBookings(movieID)

	utils.WriteJson(w, http.StatusOK, seats)
}
