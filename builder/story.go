package main

import (
	"fmt"
	"os/exec"
)

// Story is the full build definition for a Twine game.
type Story struct {
	Name string
	Deps []string
}

// ReadStoryFromFile returns a new decoded Story struct
func ReadStoryFromFile(filePath string) (*Story, error) {
	// Create Story structure
	story := &Story{}

	// Init new YAML decode
	if err := ReadYamlFile(filePath, &story); err != nil {
		return nil, err
	}

	return story, nil
}

// BuildTwineStory builds the given Story definition using Tweego
//
// This library calls out to github.com/tmedwards/tweego through os/exec.
// It assumes that the tweego binary has already been fetched (through `go get`) and added to the path.
//
// Example Tweego call:
// tweego --log-files -l \
// 				--head=analytics.html \
// 				-o dist/$game.html
// 				shared/ \
// 				games/$game
func BuildTwineStory(story *Story) error {
	output := "dist/" + story.Name + ".html"
	args := []string{
		"--log-files", "-l",
		"--head=analytics.html",
		"-o", output,
	}
	args = append(args, story.Deps...)

	cmd := exec.Command("tweego", args...)
	// fmt.Printf("cmd: %v\n", cmd)
	out, err := cmd.CombinedOutput()

	fmt.Printf("%s\n", out)
	fmt.Printf("Built '%s' as %s\n", story.Name, output)

	return err
}
