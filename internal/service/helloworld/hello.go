package helloworld

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) HelloWorld() string {
	return "hello world 2023"
}
