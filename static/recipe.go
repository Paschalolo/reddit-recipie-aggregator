package main

type Recipe struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	Ingredients []Ingredients `json:"ingredients"`
	Steps       []string      `json:"steps"`
	Picture     string        `json:"imageURL"`
}

type Ingredients struct {
	Quantity string `json:"quantity"`
	Name     string `json:"name"`
	Type     string `json:"type"`
}
