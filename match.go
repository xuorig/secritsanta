package main

import (
	"math/rand"
	"time"
)

// Match represents a match between a Secrit Santa and the gift receiver
type Match struct {
	Santa    string
	Receiver string
}

// GetMatches returns a list of matches found from the list of participant emails
func GetMatches(participants []string) []Match {
	rand.Seed(time.Now().Unix())

	var matches []Match

	alreadyPicked := map[string]bool{}

	for i, participant := range participants {
		choices := possibleChoices(i, alreadyPicked, participants)
		pick := choices[rand.Intn(len(choices))]

		alreadyPicked[pick] = true

		match := Match{
			Santa:    participant,
			Receiver: pick,
		}

		matches = append(matches, match)
	}

	return matches
}

func possibleChoices(currentIndex int, alreadyPicked map[string]bool, participants []string) []string {
	var possibleChoices []string

	for i, participant := range participants {
		if i != currentIndex && !alreadyPicked[participant] {
			possibleChoices = append(possibleChoices, participant)
		}
	}

	return possibleChoices
}
