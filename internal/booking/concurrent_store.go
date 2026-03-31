package booking

type ConcurrentStore struct {
	bookings map[string]Booking
}

func NewConcurrentStore() *ConcurrentStore {
	return &ConcurrentStore{
		bookings: map[string]Booking{},
	}
}

func (s *ConcurrentStore) Book(b Booking) error {
	if _, exists := s.bookings[b.SeatID]; exists {
		return ErrSeatAlreadyBooked
	}

	s.bookings[b.SeatID] = b
	return nil
}

func (s *ConcurrentStore) ListBookings(movieId string) []Booking {
	result := []Booking{}
	for _, b := range s.bookings {
		result = append(result, b)
	}

	return result
}
