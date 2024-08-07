package lib

import (
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/segmentio/ksuid"
	"github.com/sirupsen/logrus"
)

// Recipe is the primary model in the API.  It's use is to be stored and
// retrieved from DynamoDB.
type Recipe struct {
	CreatedTimestamp time.Time `dynamodbav:"CreatedTimestamp"`
	UpdatedTimestamp time.Time `dynamodbav:"UpdatedTimestamp"`
	ID               string    `dynamodbav:"ID"`
	PK               string    `dynamodbav:"PK"`
	SK               string    `dynamodbav:"SK"`
	Author           string    `dynamodbav:"Author"`
	Name             string    `dynamodbav:"Name"`
	Description      string    `dynamodbav:"Description"`
}

// RecipeTypesense represents the document model that will be persisted
// into the Typesense cluster
type RecipeTypesense struct {
	ID               string `json:"id"`
	Author           string `json:"author"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	CreatedTimestamp int64  `json:"createdTimestamp"`
	UpdatedTimestamp int64  `json:"updatedTimestamp"`
}

// NewRecipeTypesenseFromRecipe function for creating a RecipeTypesense
// from a Recipe
func NewRecipeTypesenseFromRecipe(recipe *Recipe) *RecipeTypesense {
	return &RecipeTypesense{
		ID:               recipe.ID,
		Author:           recipe.Author,
		Description:      recipe.Description,
		Name:             recipe.Name,
		CreatedTimestamp: recipe.CreatedTimestamp.Unix(),
		UpdatedTimestamp: recipe.UpdatedTimestamp.Unix(),
	}
}

// NewRecipeFromTypesenseRecipe function takes a map and converts it to a
// Recipe.  This is used convertind data out of Typesense
func NewRecipeFromTypesenseRecipe(m map[string]interface{}) *Recipe {
	r := &Recipe{}

	for k, v := range m {
		if k == "description" {
			r.Description = v.(string)
		} else if k == "name" {
			r.Name = v.(string)
		} else if k == "author" {
			r.Author = v.(string)
		} else if k == "createdTimestamp" {
			t := v.(float64)
			r.CreatedTimestamp = time.Unix(int64(t), 0)
		} else if k == "updatedTimestamp" {
			t := v.(float64)
			r.UpdatedTimestamp = time.Unix(int64(t), 0)
		} else if k == "id" {
			r.ID = v.(string)
		}
	}

	r.PK = fmt.Sprintf("RECIPE#%s", r.ID)
	r.SK = fmt.Sprintf("RECIPE#%s", r.ID)
	return r
}

// NewRecipeFromStreamRecord function handles converting a record from a
// DynamoDBStream Record into a Recipe
func NewRecipeFromStreamRecord(record events.DynamoDBEventRecord) *Recipe {
	r := &Recipe{}
	for k, v := range record.Change.NewImage {
		if k == "PK" {
			r.PK = v.String()
		} else if k == "SK" {
			r.SK = v.String()
		} else if k == "Description" {
			r.Description = v.String()
		} else if k == "ID" {
			r.ID = v.String()
		} else if k == "Name" {
			r.Name = v.String()
		} else if k == "Author" {
			r.Author = v.String()
		} else if k == "CreatedTimestamp" {
			t := v.String()
			createdTime, err := time.Parse("2006-01-02T15:04:05Z", t)
			if err != nil {
				logrus.Error("Error parsing CreatedTimestamp")
			} else {
				r.CreatedTimestamp = createdTime
			}
		} else if k == "UpdatedTimestamp" {
			t := v.String()
			updatedTime, err := time.Parse("2006-01-02T15:04:05Z", t)
			if err != nil {
				logrus.Error("Error parsing UpdatedTimestamp")
			} else {
				r.UpdatedTimestamp = updatedTime
			}
		}
	}
	return r
}

// NewRecipeViewsFromRecipes function takes a slice of Recipes and converts
// them into RecipeViewDtos for use in the /search API
func NewRecipeViewsFromRecipes(recipes []Recipe) []RecipeViewDto {
	r := []RecipeViewDto{}
	for _, v := range recipes {
		r = append(r, *NewRecipeViewDtoFromRecipe(v))
	}
	return r
}

// NewRecipeViewDtoFromRecipe function converts a single Recipe into a RecipeViewDto
func NewRecipeViewDtoFromRecipe(recipe Recipe) *RecipeViewDto {
	return &RecipeViewDto{
		ID:               recipe.ID,
		Name:             recipe.Name,
		Description:      recipe.Description,
		Author:           recipe.Author,
		CreatedTimestamp: recipe.CreatedTimestamp,
		UpdatedTimestamp: recipe.UpdatedTimestamp,
	}
}

// NewRecipeFromCreate builds a Recipe from the RecipeCreateDto.  This function is
// used in the POST / endpoint
func NewRecipeFromCreate(dto RecipeCreateDto) *Recipe {
	now := time.Now()
	id := ksuid.New().String()
	key := fmt.Sprintf("RECIPE:%s", id)

	return &Recipe{
		CreatedTimestamp: now,
		UpdatedTimestamp: now,
		PK:               key,
		SK:               key,
		Name:             dto.Name,
		Description:      dto.Description,
		Author:           dto.Author,
		ID:               id,
	}
}
