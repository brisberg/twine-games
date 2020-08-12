package main

import (
	"log"
	"os"
)

var gamesFile = "games/games.yml"

func main() {
	targets, err := ReadTargetsListFromFile(gamesFile)
	if err != nil {
		log.Fatalf("Failed reading targets list: %v", err)
	}

	// Create dist directory for tweego outputs
	os.MkdirAll("dist", os.ModePerm)

	// Create the index markdown file
	os.MkdirAll("genfiles/public", os.ModePerm)
	mdfile, err := os.Create("genfiles/public/README.md")
	if err != nil {
		log.Fatalf("Failed to open README output file: %v", err)
	}
	defer mdfile.Close()

	// Load index markdown template
	mdtmpl, err := loadStoryTemplate("public/README.md")
	if err != nil {
		log.Fatalf("Failed load README template: %v", err)
	}

	for _, g := range *&targets.Games {
		story, err := ReadStoryFromFile(g + "/build.yml")
		if err != nil {
			log.Printf("Failed to parse story definition %v. Error: %v", g, err)
			continue
		}

		if err = BuildTwineStory(story); err != nil {
			log.Println(err)
		}

		if err = appendToLandingMarkdown(mdfile, mdtmpl, story); err != nil {
			log.Println(err)
		}
	}
}
