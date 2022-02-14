package main

import "fmt"

// Shooter iface is a common iface to union all persons that can `shoot`
type Shooter interface {
	Shoot(target string) error
}

// Hunter is a person that may shoot in real life
type Hunter struct {
	Name  string
	Rifle string
}

// Shoot method represents how real Hunter uses rifle
func (h Hunter) Shoot(target string) error {
	fmt.Printf("Hunter %v somewhere in a forest, picking up his %v, his target is %v\n", h.Name, h.Rifle, target)
	return nil
}

// Gamer is a person who shoots virtually
type Gamer struct {
	Name  string
	Game  string
	Rifle string
}

// Shoot method represents how Gamer plays
func (g Gamer) Shoot(target string) error {
	fmt.Printf("Gamer %v in game %v, picking up his %v, his target is %v\n", g.Name, g.Game, g.Rifle, target)
	return nil
}

/* One of interface features is to
reduce code duplication by using 1 function
that takes iface type as input. In such case
that function may be called from different structures/objects/types!!
*/
// ScoreUp adds +1 score point to the person, if desired target is defeated
func ScoreUp(s Shooter, target string) error {
	err := s.Shoot(target)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	woodcutter := Hunter{
		Name:  "Bob",
		Rifle: "Winchester FOO",
	}

	suPPerSniper := Gamer{
		Name:  "DreamTeamSS$$$",
		Rifle: "AK-47",
		Game:  "Counter strike GO",
	}

	err := ScoreUp(woodcutter, "plate")
	if err != nil {
		panic(err)
	}

	err = ScoreUp(suPPerSniper, "playerTwo")
	if err != nil {
		panic(err)
	}
}
