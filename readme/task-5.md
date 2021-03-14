
TASK-5. Introduce a new RPC, that allows us to fetch a single race by its ID.

Integration TEST
----------------


1. adding rpc to racing.proto in both API  both API and racing server


  // Returns a specific Race  .
  rpc GetRace(GetRaceRequest) returns (ListRacesResponse) {
    
    //   curl http://DOMAIN_NAME/v1/races/1
    option (google.api.http) = { get: "/v1/races/{id}" };
  }

2. added the new request message. Kept the response same to reduce code duplicay 

// Request message for GetShelf method.
message GetRaceRequest {
  // The ID of the race resource to retrieve.
  int64 id = 1;
}



SERVER CHANGES
----------------

1. Added below below function to interface and implemented in racing.go in 'db' folder to get result from DB

GetRace(id int64) ([]*racing.Race, error)

2. Impleted below function to resopse the result

func (s *racingService) GetRace(ctx context.Context, in *racing.GetRaceRequest) (*racing.ListRacesResponse, error) 



REQUEST   : http://localhost:8000/v1/races/{id}
------------


curl --location --request GET 'http://localhost:8000/v1/races/1' \
--header 'Content-Type: text/plain' \
--data-raw '{"filter": { "meeting_ids": [5],"visible":false}}
'
RESPONSE
--------------------

{
    "races": [
        {
            "id": "1",
            "meetingId": "5",
            "name": "North Dakota foes",
            "number": "2",
            "visible": false,
            "advertisedStartTime": "2021-03-03T01:30:57Z"
        }
    ]
}


