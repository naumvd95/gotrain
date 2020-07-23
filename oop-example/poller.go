package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
Create poller-app for voting for colors
Structure: User ID , color mapping
1 User can’t change its decision
2 User can change decision
*/

type Person struct {
	Name           string
	SocialNetworks Sources
	DesiredColor   int
}

type Sources []Poller

type Poller interface {
	Poll(color int, poll *PollMeta) error
}

type PollMeta struct {
	Body    string
	Votings []Vote
}

type Vote struct {
	User          string
	Location      string
	ColorDesicion int
}

type Telegram struct {
	Username string
	Channel  string
}

type WhatsUpp struct {
	UserID  string
	GroupID string
}

func (s WhatsUpp) Poll(color int, poll *PollMeta) error {
	vote := Vote{
		User:          s.UserID,
		Location:      s.GroupID,
		ColorDesicion: color,
	}
	fmt.Printf("[Whatsupp] Voting: user %v choosed color %v \n", vote.User, vote.ColorDesicion)
	poll.Votings = append(poll.Votings, vote)

	return nil
}

func (s Telegram) Poll(color int, poll *PollMeta) error {
	vote := Vote{
		User:          s.Username,
		Location:      s.Channel,
		ColorDesicion: color,
	}
	fmt.Printf("[Telegram] Voting: user %v choosed color %v \n", vote.User, vote.ColorDesicion)
	poll.Votings = append(poll.Votings, vote)

	return nil
}

func getCsvFromFile(filepath string) ([]Person, error) {
	humans := []Person{}
	//Getting user-data from csv file
	file, err := os.Open(filepath)
	if err != nil {
		return humans, err
	}
	r := csv.NewReader(file)
	r.Comma = ';'
	csvFileData, err := r.ReadAll()
	if err != nil {
		return humans, err
	}

	for _, line := range csvFileData[1:] {
		sCreds := Sources{
			WhatsUpp{
				UserID:  line[1],
				GroupID: line[2],
			},
			Telegram{
				Username: line[3],
				Channel:  line[4],
			},
		}
		color, err := strconv.Atoi(line[5])
		if err != nil {
			return humans, err
		}
		humans = append(humans, Person{
			Name:           line[0],
			DesiredColor:   color,
			SocialNetworks: sCreds,
		})
	}

	return humans, nil
}

func getCsvSample() ([]Person, error) {
	humans := []Person{}
	csvData := `Name;WuName;WuGroup;TgName;TgGroup;Color
Vladislav Naumov;vnXXX1234;fakeWhatsuppGroup;@vnaumov;@csvSample;3
Bladislav Karubov;bkXXX1212w4;gpoweeer;@bkarubov;@csvSample;3
slav ubov;bkXXX1212w4;gpoweeer;@bkarubov;@csvSample;3
Blav arubov;bk212w4;gper;@bkubov;@csvSample;2
Bislav Kabov;bkXXX1212w4;gpow;@bkaruov;@csvSample;2`

	r := csv.NewReader(strings.NewReader(csvData))
	r.Comma = ';'
	csvFileData, err := r.ReadAll()
	if err != nil {
		return humans, err
	}

	for _, line := range csvFileData[1:] {
		sCreds := Sources{
			WhatsUpp{
				UserID:  line[1],
				GroupID: line[2],
			},
			Telegram{
				Username: line[3],
				Channel:  line[4],
			},
		}
		color, err := strconv.Atoi(line[5])
		if err != nil {
			return humans, err
		}
		humans = append(humans, Person{
			Name:           line[0],
			DesiredColor:   color,
			SocialNetworks: sCreds,
		})
	}

	return humans, nil

}

// sorting/getting vote results
type colorRes struct {
	Number int
	Votes  int
}

func merge(l, r []colorRes) []colorRes {
	res := make([]colorRes, 0, len(l)+len(r))

	for len(l) > 0 || len(r) > 0 {
		if len(l) == 0 {
			return append(res, r...)
		}
		if len(r) == 0 {
			return append(res, l...)
		}

		if l[0].Votes >= r[0].Votes {
			res = append(res, l[0])
			l = l[1:]
		} else {
			res = append(res, r[0])
			r = r[1:]
		}
	}
	return res
}

func mergeSort(data []colorRes) []colorRes {
	if len(data) <= 1 {
		return data
	}
	initialDivider := len(data) / 2
	var leftChunk []colorRes
	var rightChunk []colorRes
	leftChunk = mergeSort(data[:initialDivider])
	rightChunk = mergeSort(data[initialDivider:])

	return merge(leftChunk, rightChunk)
}

func getVoteRes(data *PollMeta) int {
	// key: color, value: voted times
	mappedRes := make(map[int]int)
	for _, v := range data.Votings {
		mappedRes[v.ColorDesicion]++
	}

	parsedRes := []colorRes{}
	for k, v := range mappedRes {
		parsedRes = append(parsedRes, colorRes{
			Number: k,
			Votes:  v,
		})
	}
	// get max
	sortedRes := mergeSort(parsedRes)
	if sortedRes[0].Votes == sortedRes[1].Votes {
		fmt.Println("Failed to vote for majority")
		return -1
	}
	return sortedRes[0].Number
}

func main() {
	newPoll := PollMeta{
		Body: "what color do you prefer? 1-white, 2-red, 3-blue, 4-black",
	}

	//peoples, err := getCsvSample()
	peoples, err := getCsvFromFile("pollerdata.csv")
	if err != nil {
		log.Fatal(err)
	}

	for _, human := range peoples {
		for _, net := range human.SocialNetworks {
			err := net.Poll(human.DesiredColor, &newPoll)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	fmt.Printf("Here is poll %v:\nvotings:\n %v\n\n", newPoll.Body, newPoll.Votings)

	votingResult := getVoteRes(&newPoll)
	fmt.Printf("Here is poll %v:\nresult color: %v", newPoll.Body, votingResult)
}
