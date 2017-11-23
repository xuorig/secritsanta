package main

import (
	"os"
)

func main() {
	participants := os.Args[1:]
	matches := GetMatches(participants)

	for _, match := range matches {
		SendSecritSantaMatch(match)
	}
}
