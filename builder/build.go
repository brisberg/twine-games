package main

import (
	"fmt"
	"log"
	"os/exec"
)

// readGamesList returns list of game yaml paths from a yaml file
func buildTwineStory(game *Game) error {

	out, err := exec.Command("date").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The date is %s\n", out)

	// cmd := exec.Command("go", "run", "github.com/tmedwards/tweego", "-v")
	// log.Printf("Running command and waiting for it to finish...")
	// out, err := cmd.Output()
	// log.Printf("Output: %v", out)
	// log.Printf("Command finished with error: %v", err)
	return err
}

// go run github.com/tmedwards/tweego \
//     --log-files -l \
//     --head=analytics.html \
//     shared/ \
//     games/$game \
//     -o dist/$game.html
