package lib

import (
	"time"
)

// RecipeCreateDto is the structure that holds the incoming create recipe
// request from the client
type RecipeCreateDto struct {
	Author      string `json:"author"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// RecipeViewDto is the structure that represents a view only model of the
// recipe.  Returned from the /search endpoint
type RecipeViewDto struct {
	CreatedTimestamp time.Time `json:"createdTimestamp"`
	UpdatedTimestamp time.Time `json:"updatedTimestamp"`
	ID               string    `json:"id"`
	Author           string    `json:"author"`
	Name             string    `json:"name"`
	Description      string    `json:"description"`
}
