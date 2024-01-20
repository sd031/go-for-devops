package myservice

// ExternalAPI is an interface for the external API
type ExternalAPI interface {
	FetchData(param string) (string, error)
}

// MyService uses ExternalAPI
type MyService struct {
	API ExternalAPI
}

// New creates a new instance of MyService
func New(api ExternalAPI) *MyService {
	return &MyService{API: api}
}

// DoSomething uses ExternalAPI to perform an operation
func (s *MyService) DoSomething(param string) (string, error) {
	return s.API.FetchData(param)
}
