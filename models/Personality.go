package models

type Personality struct {
	ID      string `gorm:"type:uuid;primaryKey" json:"id"`
	Name    string `json:"name"`
	History string `json:"history"`
}

var Personalities []Personality
