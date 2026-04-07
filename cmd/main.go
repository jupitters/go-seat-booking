package main

import (
	"log"
	"net/http"

	"github.com/jupitters/go-seat-booking/internal/adapters/redis"
	"github.com/jupitters/go-seat-booking/internal/booking"
	"github.com/jupitters/go-seat-booking/internal/utils"
)

type movieResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Rows        int    `json:"rows"`
	SeatsPerRow int    `json:"seats_per_row"`
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /movies", listMovies)

	mux.Handle("GET /", http.FileServer(http.Dir("static")))

	store := booking.NewRedisStore(redis.NewClient("localhost:6379"))
	svc := booking.NewService(store)

	bookinHandler := booking.NewHandler(svc)

	mux.HandleFunc("GET /movies/:movieID/seats", bookinHandler.ListSeats)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

var movies = []movieResponse{
	{ID: "inception", Title: "Inception", Rows: 5, SeatsPerRow: 8},
	{ID: "dune", Title: "Dune: Part Two", Rows: 4, SeatsPerRow: 6},
}

func listMovies(w http.ResponseWriter, r *http.Request) {
	utils.WriteJson(w, http.StatusOK, movies)
}
