package pkg

import "time"

// Recipe represents a culinary recipe.
// @Description Represents a culinary recipe with a name, ingredients, and tags.
type Recipe struct {
	// ID of the recipe.
	// @Example 123ojwfnowndno
	ID string `json:"id,omitempty"`

	// Name of the recipe.
	// @Example Delicious Pasta
	Name string `json:"name"`

	// Tags associated with the recipe.
	// @Example ["italian", "dinner"]
	Tags []string `json:"tags"`

	// List of ingredients required for the recipe.
	// @Example ["pasta", "tomato sauce", "cheese"]
	Ingredients []string `json:"ingredients"`

	// Instructions associated with the recipe.
	// @Example ["To marinate the chicken", "Add scallion whites, and cook, stirring"]
	Instructions []string `json:"instructions"`

	// Time of  the recipe added or updated .
	// @Example "2021-01-17T19:28:52.803062+01:00"
	PublishedAt time.Time `json:"publishedAt,omitempty"`
}
