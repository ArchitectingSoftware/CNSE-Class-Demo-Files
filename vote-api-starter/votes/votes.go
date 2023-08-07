package election

import (
	"encoding/json"
)

type Vote struct {
	VoteID    uint
	VoterID   uint
	PollID    uint
	VoteValue uint
}
type VoteData struct {
	Votes []Vote //A map of VoterIDs as keys and Voter structs as values
}

// constructor for VoterList struct
func NewVote(pid, vid, vtrid, vval uint) *Vote {
	return &Vote{
		VoteID:    vid,
		VoterID:   vtrid,
		PollID:    pid,
		VoteValue: vval,
	}
}

func NewSampleVote() *Vote {
	return &Vote{
		VoteID:    1,
		PollID:    1,
		VoterID:   1,
		VoteValue: 1,
	}
}

func (p *Vote) ToJson() string {
	b, _ := json.Marshal(p)
	return string(b)
}
