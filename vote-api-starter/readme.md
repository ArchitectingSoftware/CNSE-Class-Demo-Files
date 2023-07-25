## Voter API Starter

Note this is not an API, but instead some starter code for your assignment. It has the correct data structures and also seperates the management of the voter data structures from the API code.

You can thank me later :-

The output of running this code should produce:

```
➜  vote-api-starter git:(main) ✗ go run main.go
------------------------
{"VoterID":0,"FirstName":"John","LastName":"Doe","VoteHistory":[{"PollID":1,"VoteDate":"2023-07-25T19:10:34.811997-04:00"},{"PollID":2,"VoteDate":"2023-07-25T19:10:34.811998-04:00"}]}
------------------------
{"VoterID":0,"FirstName":"Jane","LastName":"Doe","VoteHistory":[{"PollID":1,"VoteDate":"2023-07-25T19:10:34.811998-04:00"},{"PollID":2,"VoteDate":"2023-07-25T19:10:34.811998-04:00"}]}
------------------------
{"Voters":{"1":{"VoterID":0,"FirstName":"John","LastName":"Doe","VoteHistory":[{"PollID":1,"VoteDate":"2023-07-25T19:10:34.811997-04:00"},{"PollID":2,"VoteDate":"2023-07-25T19:10:34.811998-04:00"}]},"2":{"VoterID":0,"FirstName":"Jane","LastName":"Doe","VoteHistory":[{"PollID":1,"VoteDate":"2023-07-25T19:10:34.811998-04:00"},{"PollID":2,"VoteDate":"2023-07-25T19:10:34.811998-04:00"}]}}}
------------------------
```