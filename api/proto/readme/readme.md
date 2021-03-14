Code changes and manual integration test
------------------------------------

TASK 3: 3. Our races require a new `status` field that is derived based on their `advertised_start_time`'s. The status is simply, `OPEN` or `CLOSED`. All races that have an `advertised_start_time` in the passed should reflect `CLOSED`.

steps
-----------------------
 RACE Server side
------------------

1.  Update query in getRaceQueries

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

2. Update below line races.go

if err := rows.Scan(&race.Id, &race.MeetingId, &race.Name, &race.Number, &race.Visible, &advertisedStart, &race.Status);


3. Add  status field to send in response.   in File racing.proto

// status repressent whether race status is OPEN or CLOSED
  string status = 7;

4. execute below command to generage RPC files

   protoc -IC:\santhosh\protoc-3.13.0-win64\include -I . --go_out=. --go-grpc_out=require_unimplemented_servers=false:. racing/racing.proto


5. start server

   go build && racing



-----------------------
 API Client side
------------------

1. Add status to Racing to see status in response.

string status = 7;

2.  execute below command to generate RPC files

protoc -IC:\santhosh\protoc-3.13.0-win64\include -I%GOPATH%\src -I . --go_out . --go_opt paths=source_relative --go-grpc_out . --go-grpc_opt paths=source_relative --grpc-gateway_out . --grpc-gateway_opt paths=source_relative racing/racing.proto


3. start API

go build && api


Request:

curl --location --request POST 'http://localhost:8000/v1/list-races' \
--header 'Content-Type: text/plain' \
--data-raw '{"filter": { "meeting_ids": [5],"visible":false}}
'

Result:


{"races":[{"id":"1", "meetingId":"5", "name":"North Dakota foes", "number":"2", "visible":false, "advertisedStartTime":"2021-03-03T01:30:57Z", "status":"CLOSED"}, {"id":"14", "meetingId":"5", "name":"Connecticut sorcerors", "number":"4", "visible":true, "advertisedStartTime":"2021-03-03T02:41:45Z", "status":"CLOSED"}, {"id":"15", "meetingId":"5", "name":"Washington rabbits", "number":"12", "visible":false, "advertisedStartTime":"2021-03-02T02:01:14Z", "status":"CLOSED"}, {"id":"20", "meetingId":"5", "name":"Washington giants", "number":"12", "visible":false, "advertisedStartTime":"2021-03-01T12:55:50Z", "status":"CLOSED"}, {"id":"31", "meetingId":"5", "name":"Kentucky chickens", "number":"2", "visible":false, "advertisedStartTime":"2021-03-02T11:46:46Z", "status":"CLOSED"}, {"id":"33", "meetingId":"5", "name":"Kentucky black cats", "number":"4", "visible":false, "advertisedStartTime":"2021-03-01T13:41:26Z", "status":"CLOSED"}, {"id":"45", "meetingId":"5", "name":"New York cats", "number":"2", "visible":true, "advertisedStartTime":"2021-03-01T00:20:31Z", "status":"CLOSED"}, {"id":"51", "meetingId":"5", "name":"New Mexico spirits", "number":"5", "visible":false, "advertisedStartTime":"2021-02-28T17:47:56Z", "status":"CLOSED"}, {"id":"58", "meetingId":"5", "name":"New York cattle", "number":"7", "visible":false, "advertisedStartTime":"2021-02-28T20:50:39Z", "status":"CLOSED"}, {"id":"66", "meetingId":"5", "name":"Illinois sorcerors", "number":"1", "visible":true, "advertisedStartTime":"2021-03-03T05:28:20Z", "status":"CLOSED"}, {"id":"68", "meetingId":"5", "name":"Maryland warlocks", "number":"2", "visible":true, "advertisedStartTime":"2021-03-03T03:42:53Z", "status":"CLOSED"}, {"id":"78", "meetingId":"5", "name":"Rhode Island vixens", "number":"11", "visible":true, "advertisedStartTime":"2021-03-03T02:29:40Z", "status":"CLOSED"}, {"id":"87", "meetingId":"5", "name":"Wyoming frogs", "number":"9", "visible":false, "advertisedStartTime":"2021-02-28T20:15:16Z", "status":"CLOSED"}, {"id":"91", "meetingId":"5", "name":"Missouri bees", "number":"1", "visible":false, "advertisedStartTime":"2021-03-02T04:18:16Z", "status":"CLOSED"}, {"id":"97", "meetingId":"5", "name":"Arkansas geese", "number":"2", "visible":false, "advertisedStartTime":"2021-02-28T22:02:34Z", "status":"CLOSED"}]}
