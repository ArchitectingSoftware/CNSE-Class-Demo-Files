package election

import (
	"encoding/json"
)

type Vote struct {
	PollID    uint
	VoterID   uint
	VoteValue uint
}
type VoteData struct {
	Votes []Vote //A map of VoterIDs as keys and Voter structs as values
}

// constructor for VoterList struct
func NewVote(pid, vid, vval uint) *Vote {
	return &Vote{
		PollID:    pid,
		VoterID:   vid,
		VoteValue: vval,
	}
}

func NewSampleVote() *Vote {
	return &Vote{
		PollID:    1,
		VoterID:   1,
		VoteValue: 1,
	}
}

func (p *Vote) ToJson() string {
	b, _ := json.Marshal(p)
	return string(b)
}
