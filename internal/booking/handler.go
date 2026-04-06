package booking

type handler struct {
	svc Service
}

func NewHandler() *handler {
	return &handler{}
}
