package booking

type ConcurrentStore struct {
	bookings map[string]Booking
}

func NewConcurrentStore() *MemoryStore {
	return &MemoryStore{
		bookings: map[string]Booking{},
	}
}

func (s *MemoryStore) Book(b Booking) error {
	if _, exists := s.bookings[b.SeatID]; exists {
		return ErrSeatAlreadyBooked
	}

	s.bookings[b.SeatID] = b
	return nil
}

func (s *MemoryStore) ListBookings(movieId string) []Booking {
	result := []Booking{}
	for _, b := range s.bookings {
		result = append(result, b)
	}

	return result
}
