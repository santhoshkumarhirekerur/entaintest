package db

// import
import (
	"database/sql"
	"strings"
	"sync"
	"time"

	"git.neds.sh/matty/entain/racing/proto/sports"

	"github.com/golang/protobuf/ptypes"
	_ "github.com/mattn/go-sqlite3"
)

// SportsRepo provides repository access to sports event.
type SportsRepo interface {
	// Init will initialise our races repository.
	Init() error

	// List will return a list of races.
	List(filter *sports.ListSportsEventRequestFilter) ([]*sports.SportEvent, error)
}

type sportsRepo struct {
	db   *sql.DB
	init sync.Once
}

// NewSportsRepo creates a new sports repository.
func NewSportsRepo(db *sql.DB) SportsRepo {
	return &sportsRepo{db: db}
}

// Init prepares the race repository dummy data.
func (s *sportsRepo) Init() error {
	var err error

	s.init.Do(func() {
		// For test/example purposes, we seed the DB with some dummy races.
		err = s.seed()
	})

	return err
}

func (s *sportsRepo) List(filter *sports.ListSportsEventRequestFilter) ([]*sports.SportEvent, error) {
	var (
		err   error
		query string
		args  []interface{}
	)

	query = getSportsEventQueries()[sportsEventList]

	query, args = s.applyFilter(query, filter)

	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	return s.scanSportEvents(rows)
}

func (r *sportsRepo) applyFilter(query string, filter *sports.ListSportsEventRequestFilter) (string, []interface{}) {
	var (
		clauses []string
		args    []interface{}
	)

	if filter == nil {
		return query, args
	}

	if len(filter.Names) > 0 {
		clauses = append(clauses, "name IN ("+strings.Repeat("?,", len(filter.Names)-1)+"?)")

		for _, name := range filter.Names {
			args = append(args, name)
		}
	}

	if len(clauses) != 0 {
		query += " WHERE " + strings.Join(clauses, " AND ")
	}

	return query, args
}

func (m *sportsRepo) scanSportEvents(
	rows *sql.Rows,
) ([]*sports.SportEvent, error) {
	var sportsEvents []*sports.SportEvent

	for rows.Next() {
		var sportEvent sports.SportEvent
		var advertisedStart time.Time

		if err := rows.Scan(&sportEvent.Id, &sportEvent.Name, &advertisedStart); err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}

			return nil, err
		}

		ts, err := ptypes.TimestampProto(advertisedStart)
		if err != nil {
			return nil, err
		}

		sportEvent.AdvertisedStartTime = ts

		sportsEvents = append(sportsEvents, &sportEvent)
	}

	return sportsEvents, nil
}
