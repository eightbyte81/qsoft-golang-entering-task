package service

type TimeDifference interface {
	GetTimeDifference(yearString string) (float64, error)
}

type Service struct {
	TimeDifference
}

func NewService() *Service {
	return &Service{
		TimeDifference: NewTimeDifferenceService(),
	}
}
