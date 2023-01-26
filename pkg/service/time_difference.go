package service

import (
	"strconv"
	"time"
)

type TimeDifferenceService struct{}

func NewTimeDifferenceService() *TimeDifferenceService {
	return &TimeDifferenceService{}
}

func (s *TimeDifferenceService) GetTimeDifference(yearString string) (float64, error) {
	year, err := strconv.Atoi(yearString)
	if err != nil {
		return 0, err
	}

	reqTime := time.Date(year, 1, 1, 0, 0, 0, 0, time.Local)

	return time.Now().Sub(reqTime).Hours() / 24, nil
}
