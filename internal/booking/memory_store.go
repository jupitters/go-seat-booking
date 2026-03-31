package booking

type MemoryStore struct {
	bookings map[string]Booking
}

func NewMemoryStore() *MemoryStore {
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
