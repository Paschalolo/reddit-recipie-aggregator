package pkg

import (
	"time"
)

// Recipe represents a culinary recipe.
// @Description Represents a culinary recipe with a name, ingredients, and tags.
type Recipe struct {
	// ID of the recipe.
	// @Example 123ojwfnowndno
	ID string `json:"id" bson:"_id"`

	// Name of the recipe.
	// @Example Delicious Pasta
	Name string `json:"name,omitempty" bson:"name"`

	// Tags associated with the recipe.
	// @Example ["italian", "dinner"]
	Tags []string `json:"tags,omitempty" bson:"tags"`

	// List of ingredients required for the recipe.
	// @Example ["pasta", "tomato sauce", "cheese"]
	Ingredients []string `json:"ingredients,omitempty" bson:"ingredients"`

	// Instructions associated with the recipe.
	// @Example ["To marinate the chicken", "Add scallion whites, and cook, stirring"]
	Instructions []string `json:"instructions,omitempty" bson:"instructions"`

	// Time of  the recipe added or updated .
	// @Example "2021-01-17T19:28:52.803062+01:00"
	PublishedAt time.Time `json:"publishedAt,omitempty" bson:"publishedAt"`
}
