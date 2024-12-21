package models

type Personality struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	History string `json:"history"`
}

var Personalities []Personality
