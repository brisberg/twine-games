package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

var gamesFile = "games/games.yml"

// gamesList is the yaml definition a list of game definitions
type gamesList struct {
	Games []string
}

// readGamesList returns list of game yaml paths from a yaml file
func readGamesList(path string) (*[]string, error) {
	// Create games list
	list := &gamesList{}

	// Open game file
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&list); err != nil {
		return nil, err
	}

	return &list.Games, nil
}

func main() {
	list, err := readGamesList(gamesFile)
	if err != nil {
		log.Fatalf("error: %v", err)
		os.Exit(1)
	}
	fmt.Printf("--- m dump:\n%s\n\n", list)

	for i, g := range *list {
		game, err := ReadGame(g + "/build.yml")
		if err != nil {
			log.Printf("Failed to parse game definition %v. Error: %v", g, err)
			continue
		}

		d, err := yaml.Marshal(&game)
		fmt.Printf("--- m dump:\n%s\n\n", string(d))
		fmt.Println(i, g)
	}

	buildTwineStory(&Game{})
}
