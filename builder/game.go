package main

import (
	"os"

	"gopkg.in/yaml.v2"
)

// Note: struct fields must be public in order for unmarshal to
// correctly populate the data.

// Game is the full build definition for a Twine game.
type Game struct {
	Name string
	Deps []string
}

// ReadGame returns a new decoded Game struct
func ReadGame(gamePath string) (*Game, error) {
	// Create game structure
	game := &Game{}

	// Open game file
	file, err := os.Open(gamePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&game); err != nil {
		return nil, err
	}

	return game, nil
}
