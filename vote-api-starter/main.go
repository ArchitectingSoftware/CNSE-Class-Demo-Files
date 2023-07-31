package main

import (
	"fmt"
	"voter-api-starter/api"
)

func main() {
	//v := voter.NewVoter(1, "John", "Doe")
	//v.AddPoll(1)
	//v.AddPoll(2)
	//v.AddPoll(3)
	//v.AddPoll(4)

	//b, _ := json.Marshal(v)
	//fmt.Println(string(b))
	vl := api.NewVoterApi()

	//POST /voter with body of John Doe
	vl.AddVoter(1, "John", "Doe")

	//POST /voter/1/poll with body of a voter record only including the Poll information
	vl.LetsSimulateAPostForAPoll(1)
	fmt.Println()
	vl.LetsSimulateAPostForAPoll(5)

	/*
		vl.AddVoter(1, "John", "Doe")
		vl.AddPoll(1, 1)
		vl.AddPoll(1, 2)
		vl.AddVoter(2, "Jane", "Doe")
		vl.AddPoll(2, 1)
		vl.AddPoll(2, 2)

		fmt.Println("------------------------")
		fmt.Println(vl.GetVoterJson(1))
		fmt.Println("------------------------")
		fmt.Println(vl.GetVoterJson(2))
		fmt.Println("------------------------")
		fmt.Println(vl.GetVoterListJson())
		fmt.Println("------------------------")
	*/
}
