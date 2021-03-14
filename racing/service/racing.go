package service

import (
	"git.neds.sh/matty/entain/racing/db"
	"git.neds.sh/matty/entain/racing/proto/racing"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

type Racing interface {
	// ListRaces will return a collection of races.
	ListRaces(ctx context.Context, in *racing.ListRacesRequest) (*racing.ListRacesResponse, error)
}

// racingService implements the Racing interface.
type racingService struct {
	racesRepo db.RacesRepo
}

// NewRacingService instantiates and returns a new racingService.
func NewRacingService(racesRepo db.RacesRepo) Racing {
	return &racingService{racesRepo}
}

func (s *racingService) ListRaces(ctx context.Context, in *racing.ListRacesRequest) (*racing.ListRacesResponse, error) {

	log.Infof("GETVIEW: %d", in.GetView())

	if in.GetView() == 1 {
		races, err := s.racesRepo.ListView(in.Filter)
		if err != nil {
			return nil, err
		}
		return &racing.ListRacesResponse{Races: races}, nil

	} else {
		races, err := s.racesRepo.List(in.Filter)
		if err != nil {
			return nil, err
		}
		return &racing.ListRacesResponse{Races: races}, nil
	}

}
