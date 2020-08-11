package main

import (
	"log"
	"os"
)

var gamesFile = "games/games.yml"

func main() {
	targets, err := ReadTargetsListFromFile(gamesFile)
	if err != nil {
		log.Fatalf("error: %v", err)
		os.Exit(1)
	}

	for _, g := range *&targets.Games {
		game, err := ReadStoryFromFile(g + "/build.yml")
		if err != nil {
			log.Printf("Failed to parse game definition %v. Error: %v", g, err)
			continue
		}

		BuildTwineStory(game)
	}
}
