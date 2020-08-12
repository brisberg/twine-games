package main

import (
	"log"
)

var gamesFile = "games/games.yml"

func main() {
	targets, err := ReadTargetsListFromFile(gamesFile)
	if err != nil {
		log.Fatalf("Failed reading targets list: %v", err)
	}

	for _, g := range *&targets.Games {
		game, err := ReadStoryFromFile(g + "/build.yml")
		if err != nil {
			log.Printf("Failed to parse game definition %v. Error: %v", g, err)
			continue
		}

		if err = BuildTwineStory(game); err != nil {
			log.Println(err)
		}
	}
}
