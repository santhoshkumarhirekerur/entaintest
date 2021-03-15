package service

import (
	"git.neds.sh/matty/entain/racing/proto/sports"

	"git.neds.sh/matty/entain/racing/db"

	"golang.org/x/net/context"
)

type Sports interface {
	// ListRaces will return a collection of races.
	ListSportEvent(ctx context.Context, in *sports.ListSportsEventRequest) (*sports.ListSportsEventResponse, error)
}

// racingService implements the Racing interface.
type sportsService struct {
	sportsRepo db.SportsRepo
}

// NewRacingService instantiates and returns a new racingService.
func NewSportsService(sportsRepo db.SportsRepo) Sports {
	return &sportsService{sportsRepo}
}

func (s *sportsService) ListSportEvent(ctx context.Context, in *sports.ListSportsEventRequest) (*sports.ListSportsEventResponse, error) {
	sportEvents, err := s.sportsRepo.List(in.Filter)
	if err != nil {
		return nil, err
	}

	return &sports.ListSportsEventResponse{Sportevents: sportEvents}, nil
}
