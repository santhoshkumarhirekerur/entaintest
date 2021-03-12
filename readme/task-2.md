1. add conditon order by ORDER by advertised_start_time

query += "ORDER by advertised_start_time"

2. curl request

curl --location --request POST 'http://localhost:8000/v1/list-races' \
--header 'Content-Type: text/plain' \
--data-raw '{"filter": { "meeting_ids": [5],"visible":true}}
'

3. resonse  is below. dates are sorted in acesending order

{"races":[{"id":"51","meetingId":"5","name":"New Mexico spirits","number":"5","visible":false,"advertisedStartTime":"2021-02-28T17:47:56Z"},{"id":"87","meetingId":"5","name":"Wyoming frogs","number":"9","visible":false,"advertisedStartTime":"2021-02-28T20:15:16Z"},{"id":"58","meetingId":"5","name":"New York cattle","number":"7","visible":false,"advertisedStartTime":"2021-02-28T20:50:39Z"},{"id":"97","meetingId":"5","name":"Arkansas geese","number":"2","visible":false,"advertisedStartTime":"2021-02-28T22:02:34Z"},{"id":"45","meetingId":"5","name":"New York cats","number":"2","visible":true,"advertisedStartTime":"2021-03-01T00:20:31Z"},{"id":"20","meetingId":"5","name":"Washington giants","number":"12","visible":false,"advertisedStartTime":"2021-03-01T12:55:50Z"},{"id":"33","meetingId":"5","name":"Kentucky black cats","number":"4","visible":false,"advertisedStartTime":"2021-03-01T13:41:26Z"},{"id":"15","meetingId":"5","name":"Washington rabbits","number":"12","visible":false,"advertisedStartTime":"2021-03-02T02:01:14Z"},{"id":"91","meetingId":"5","name":"Missouri bees","number":"1","visible":false,"advertisedStartTime":"2021-03-02T04:18:16Z"},{"id":"31","meetingId":"5","name":"Kentucky chickens","number":"2","visible":false,"advertisedStartTime":"2021-03-02T11:46:46Z"},{"id":"1","meetingId":"5","name":"North Dakota foes","number":"2","visible":false,"advertisedStartTime":"2021-03-03T01:30:57Z"},{"id":"78","meetingId":"5","name":"Rhode Island vixens","number":"11","visible":true,"advertisedStartTime":"2021-03-03T02:29:40Z"},{"id":"14","meetingId":"5","name":"Connecticut sorcerors","number":"4","visible":true,"advertisedStartTime":"2021-03-03T02:41:45Z"},{"id":"68","meetingId":"5","name":"Maryland warlocks","number":"2","visible":true,"advertisedStartTime":"2021-03-03T03:42:53Z"},{"id":"66","meetingId":"5","name":"Illinois sorcerors","number":"1","visible":true,"advertisedStartTime":"2021-03-03T05:28:20Z"}]}