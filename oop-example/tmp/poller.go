package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

/*
Create poller-app for voting for colors
Structure: User ID , color mapping
1 User can’t change its decision
2 User can change decision
*/

type Person struct {
	Name           string
	DesiredColor   int
	SocialNetworks Sources
}
type Sources []Poller

type PollData struct {
	Body    string
	Votings []Vote
	Result  int
}

type Vote struct {
	User         string
	Group        string
	ChoosedColor int
}

type Poller interface {
	Poll(color int, poll *PollData) error
}

type Telegram struct {
	Username string
	Channel  string
}

type Viber struct {
	UserID  string
	GroupID string
}

func (s Telegram) Poll(color int, poll *PollData) error {
	vote := Vote{
		User:         s.Username,
		Group:        s.Channel,
		ChoosedColor: color,
	}
	fmt.Printf("[TELEGRAM]: User %v in group %v voted for %v\n", vote.User, vote.Group, vote.ChoosedColor)
	poll.Votings = append(poll.Votings, vote)
	return nil
}

func (s Viber) Poll(color int, poll *PollData) error {
	vote := Vote{
		User:         s.UserID,
		Group:        s.GroupID,
		ChoosedColor: color,
	}
	fmt.Printf("[VIBER]: User %v in group %v voted for %v\n", vote.User, vote.Group, vote.ChoosedColor)
	poll.Votings = append(poll.Votings, vote)
	return nil
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
		if l[0].VotedTimes >= r[0].VotedTimes {
			res = append(res, l[0])
			l = l[1:]
		} else {
			res = append(res, r[0])
			r = r[1:]
		}
	}

	return res
}

var sem = make(chan struct{}, 100)

func mergeSortMulti(data []colorRes) []colorRes {
	if len(data) <= 1 {
		return data
	}

	initialDivider := len(data) / 2
	var l []colorRes
	var r []colorRes

	wg := sync.WaitGroup{}
	wg.Add(2)
	select {
	case sem <- struct{}{}:
		go func() {
			l = mergeSortMulti(data[:initialDivider])
			wg.Done()
		}()
	default:
		l = mergeSortMulti(data[:initialDivider])
		wg.Done()
	}

	select {
	case sem <- struct{}{}:
		go func() {
			r = mergeSortMulti(data[initialDivider:])
			wg.Done()
		}()
	default:
		r = mergeSortMulti(data[initialDivider:])
		wg.Done()
	}
	wg.Wait()

	return merge(l, r)
}

type colorRes struct {
	ColorNumber int
	VotedTimes  int
}

func mapVote(data *PollData) []colorRes {
	res := []colorRes{}
	dubMap := make(map[int]int)

	for _, v := range data.Votings {
		dubMap[v.ChoosedColor]++
	}
	for k, v := range dubMap {
		res = append(res, colorRes{
			ColorNumber: k,
			VotedTimes:  v,
		})
	}

	return res
}

func main() {
	file, err := os.Open("pollerdata.csv")
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(file)
	r.Comma = ';'
	lines, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	peoples := []Person{}
	for _, set := range lines[1:] {
		color, err := strconv.Atoi(set[5])
		if err != nil {
			log.Fatal(err)
		}

		peoples = append(peoples, Person{
			Name: set[0],
			SocialNetworks: Sources{
				Viber{
					UserID:  set[1],
					GroupID: set[2],
				},
				Telegram{
					Username: set[3],
					Channel:  set[4],
				},
			},
			DesiredColor: color,
		})
	}
	poll := PollData{
		Body: "Which color do u prefer?",
	}

	for _, human := range peoples {
		for _, net := range human.SocialNetworks {
			err := net.Poll(human.DesiredColor, &poll)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	mappedRes := mapVote(&poll)
	fmt.Println(mappedRes)
	sortedRes := mergeSortMulti(mappedRes)
	fmt.Printf("Sorted result: %v\n", sortedRes[0].ColorNumber)

}
