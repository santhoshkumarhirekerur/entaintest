steps
-------
--------------------------------
API client changes
---------------------------------
1. set up go on Windows
2. set up grpc on Windows
3.  add visible field to filter api\proto\racing\racing.proto

// Filter for listing races.
message ListRacesRequestFilter {
  repeated int64 meeting_ids = 1;
  bool visible = 2;
}

4. execute below command from proto folder to auto generage racing_grpc.pb.go, racing.pb.gw.go and racing.pb.go

protoc -IC:\santhosh\protoc-3.13.0-win64\include -I%GOPATH%\src -I . --go_out . --go_opt paths=source_relative --go-grpc_out . --go-grpc_opt paths=source_relative --grpc-gateway_out . --grpc-gateway_opt paths=source_relative racing/racing.proto

5. execute below command from api folder
go build && api

--------------------------------------------------
RACE Server changes
--------------------------------------------------
1. add visible field to filter racing\proto\racing\racing.pb.go
 bool visible = 2;

2. update query condition in file db/races.go by adding below line. This is for visible filter

clauses = append(clauses, "visible=?")
args = append(args, filter.Visible)


3. execute below command

protoc -IC:\santhosh\protoc-3.13.0-win64\include -I . --go_out=. --go-grpc_out=require_unimplemented_servers=false:. racing/racing.proto

4. execute go build && racing

5. test from postman or cmd
 
 request for visible false
 ------

curl --location --request POST 'http://localhost:8000/v1/list-races' \
--header 'Content-Type: text/plain' \
--data-raw '{"filter": { "meeting_ids": [5],"visible":false}}

response
-----

{"races":[{"id":"1", "meetingId":"5", "name":"North Dakota foes", "number":"2", "visible":false, "advertisedStartTime":"2021-03-03T01:30:57Z"}, {"id":"15", "meetingId":"5", "name":"Washington rabbits", "number":"12", "visible":false, "advertisedStartTime":"2021-03-02T02:01:14Z"}, {"id":"20", "meetingId":"5", "name":"Washington giants", "number":"12", "visible":false, "advertisedStartTime":"2021-03-01T12:55:50Z"}, {"id":"31", "meetingId":"5", "name":"Kentucky chickens", "number":"2", "visible":false, "advertisedStartTime":"2021-03-02T11:46:46Z"}, {"id":"33", "meetingId":"5", "name":"Kentucky black cats", "number":"4", "visible":false, "advertisedStartTime":"2021-03-01T13:41:26Z"}, {"id":"51", "meetingId":"5", "name":"New Mexico spirits", "number":"5", "visible":false, "advertisedStartTime":"2021-02-28T17:47:56Z"}, {"id":"58", "meetingId":"5", "name":"New York cattle", "number":"7", "visible":false, "advertisedStartTime":"2021-02-28T20:50:39Z"}, {"id":"87", "meetingId":"5", "name":"Wyoming frogs", "number":"9", "visible":false, "advertisedStartTime":"2021-02-28T20:15:16Z"}, {"id":"91", "meetingId":"5", "name":"Missouri bees", "number":"1", "visible":false, "advertisedStartTime":"2021-03-02T04:18:16Z"}, {"id":"97", "meetingId":"5", "name":"Arkansas geese", "number":"2", "visible":false, "advertisedStartTime":"2021-02-28T22:02:34Z"}]}




request for visible true 
 ------
curl --location --request POST 'http://localhost:8000/v1/list-races' \
--header 'Content-Type: text/plain' \
--data-raw '{"filter": { "meeting_ids": [5],"visible":true}}

respose
----------------

{"races":[{"id":"14", "meetingId":"5", "name":"Connecticut sorcerors", "number":"4", "visible":true, "advertisedStartTime":"2021-03-03T02:41:45Z"}, {"id":"45", "meetingId":"5", "name":"New York cats", "number":"2", "visible":true, "advertisedStartTime":"2021-03-01T00:20:31Z"}, {"id":"66", "meetingId":"5", "name":"Illinois sorcerors", "number":"1", "visible":true, "advertisedStartTime":"2021-03-03T05:28:20Z"}, {"id":"68", "meetingId":"5", "name":"Maryland warlocks", "number":"2", "visible":true, "advertisedStartTime":"2021-03-03T03:42:53Z"}, {"id":"78", "meetingId":"5", "name":"Rhode Island vixens", "number":"11", "visible":true, "advertisedStartTime":"2021-03-03T02:29:40Z"}]}
