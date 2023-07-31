package poll

import (
	"encoding/json"
)

type pollOption struct {
	PollOptionID    uint
	PollOptionValue string
}

type Poll struct {
	PollID       uint
	PollTitle    string
	PollQuestion string
	PollOptions  []pollOption
}
type PollList struct {
	Polls map[uint]Poll //A map of VoterIDs as keys and Voter structs as values
}

// constructor for VoterList struct
func NewPoll(id uint, title, question string) *Poll {
	return &Poll{
		PollID:       id,
		PollTitle:    title,
		PollQuestion: question,
		PollOptions:  []pollOption{},
	}
}

func NewSamplePoll() *Poll {
	return &Poll{
		PollID:       1,
		PollTitle:    "Favorite Pet",
		PollQuestion: "What type of pet do you like best?",
		PollOptions: []pollOption{
			{PollOptionID: 1, PollOptionValue: "Dog"},
			{PollOptionID: 2, PollOptionValue: "Cat"},
			{PollOptionID: 3, PollOptionValue: "Fish"},
			{PollOptionID: 4, PollOptionValue: "Bird"},
			{PollOptionID: 5, PollOptionValue: "NONE"},
		},
	}
}

func (p *Poll) ToJson() string {
	b, _ := json.Marshal(p)
	return string(b)
}
