package api

import (
	"encoding/json"
	"fmt"
	"voter-api-starter/voter"
)

type VoterApi struct {
	voterList voter.VoterList
}

func NewVoterApi() *VoterApi {
	return &VoterApi{
		voterList: voter.VoterList{
			Voters: make(map[uint]voter.Voter),
		},
	}
}

func (v *VoterApi) AddVoter(voterID uint, firstName, lastName string) {
	v.voterList.Voters[voterID] = *voter.NewVoter(voterID, firstName, lastName)
}

func (v *VoterApi) AddPoll(voterID, pollID uint) {
	voter := v.voterList.Voters[voterID]
	voter.AddPoll(pollID)
	v.voterList.Voters[voterID] = voter
}

func (v *VoterApi) GetVoter(voterID uint) voter.Voter {
	voter := v.voterList.Voters[voterID]
	return voter
}

func (v *VoterApi) GetVoterJson(voterID uint) string {
	voter := v.voterList.Voters[voterID]
	return voter.ToJson()
}

func (v *VoterApi) GetVoterList() voter.VoterList {
	return v.voterList
}

func (v *VoterApi) GetVoterListJson() string {
	b, _ := json.Marshal(v.voterList)
	return string(b)
}

func (v *VoterApi) LetsSimulateAPostForAPoll(pollID uint) {

	//1.  Extract the JSON From the POST body.  In Gin you can do this with
	// ShouldBindJson helper.  But its magic, you can see this by going to
	// c.Request.Body

	jsonString := `
	{
		"VoterID": 1,
		"VoteHistory": [ {
			"PollId": ` + fmt.Sprint(pollID) + `,
			"VoteDate": "2018-09-22T12:42:31Z"
		} ]
	 }
	`
	var voterNewPoll voter.Voter
	json.Unmarshal([]byte(jsonString), &voterNewPoll)

	//2. Add the poll to the voter

	//3. Look up the voter from the voter database being managed by the API
	voter := v.voterList.Voters[voterNewPoll.VoterID]
	fmt.Println(voter)

	//4 Add the poll from the body to the voter, using the receiver in the voter
	voter.AddPollWithTimeDetails(voterNewPoll.VoteHistory[0].PollID,
		voterNewPoll.VoteHistory[0].VoteDate)
	v.voterList.Voters[voterNewPoll.VoterID] = voter
	fmt.Println("------------------------")
	fmt.Println(voter)

}
