package db

const (
	racesList       = "list"
	sportsEventList = "list"
)

func getRaceQueries() map[string]string {
	return map[string]string{
		racesList: `
			SELECT 
				id, 
				meeting_id, 
				name, 
				number, 
				visible, 
				advertised_start_time 
			FROM races
		`,
	}
}

func getSportsEventQueries() map[string]string {
	return map[string]string{
		sportsEventList: `
			SELECT 
				id, 
				name, 
				advertised_start_time 
			FROM sportevents
		`,
	}
}
