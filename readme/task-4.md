1. API change
-------------------

1. Adding RaceView


2. Racing server chages
---------------------


REQUEST
----------------

curl --location --request GET 'http://localhost:8000/v1/list-races?view=BASIC' \
--header 'Content-Type: text/plain' \
--data-raw '{"filter": { "meeting_ids": [5],"visible":false}}
'


RESPONSE
--------------


"races": [
        {
            "name": "North Dakota foes",
            "number": "2",
            "raceFull": null
        },
        {
            "name": "Connecticut griffins",
            "number": "12",
            "raceFull": null
        },
        {
            "name": "Rhode Island ghosts",
            "number": "3",
            "raceFull": null
        },
]


Note: I am new to Golang and learned go lang in last 3 days and try my best to resolve tasks with expected output. Looks this is bit advance concept. Not finding right google good example other than [Resource Views](https://cloud.google.com/apis/design/design_patterns#resource_view). This I can learn by looking into existing code in entain! :).
   
I spent extensive time to understand and try my best to get expected ouput. Not able to complete this task 100%. 

