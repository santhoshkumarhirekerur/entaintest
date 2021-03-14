package db

const (
	racesList = "list"
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
				advertised_start_time, (case when advertised_start_time > date('now')  then 'OPEN' else 'CLOSED' end) as status
			FROM races
		`,
	}
}
