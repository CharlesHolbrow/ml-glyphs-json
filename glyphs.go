package main

import (
	"encoding/json"
	"os"
)

// Glyphs is a map collection of glyphs by name string
type Glyphs map[string][]byte

// ReadGlyphsDotJSON a file 'glyphs.json` and return a collectino of glyphs
func ReadGlyphsDotJSON() (Glyphs, error) {
	glyphs := make(Glyphs, 32)

	file, err := os.Open("glyphs.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&glyphs)
	if err != nil {
		return nil, err
	}

	return glyphs, nil
}
