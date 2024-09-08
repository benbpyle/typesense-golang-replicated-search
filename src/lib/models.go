package lib

import (
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/sirupsen/logrus"
)

// Recipe is the primary model in the API.  It's use is to be stored and
// retrieved from DynamoDB.
type Todo struct {
	CreatedAt   time.Time `dynamodbav:"created_at"`
	UpdatedAt   time.Time `dynamodbav:"updated_at"`
	Id          string    `dynamodbav:id"`
	Name        string    `dynamodbav:"name"`
	Description string    `dynamodbav:"description"`
}

// RecipeTypesense represents the document model that will be persisted
// into the Typesense cluster
type TodoTypesense struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

// NewRecipeTypesenseFromRecipe function for creating a RecipeTypesense
// from a Recipe
func NewTodoTypesenseFromTodo(todo *Todo) *TodoTypesense {
	return &TodoTypesense{
		ID:          todo.Id,
		Description: todo.Description,
		Name:        todo.Name,
		CreatedAt:   todo.CreatedAt.Unix(),
		UpdatedAt:   todo.UpdatedAt.Unix(),
	}
}

// NewRecipeFromTypesenseRecipe function takes a map and converts it to a
// Recipe.  This is used convertind data out of Typesense
func NewTodoFromTypesenseTodo(m map[string]interface{}) *Todo {
	r := &Todo{}

	for k, v := range m {
		if k == "description" {
			r.Description = v.(string)
		} else if k == "name" {
			r.Name = v.(string)
		} else if k == "created_at" {
			t := v.(float64)
			r.CreatedAt = time.Unix(int64(t), 0)
		} else if k == "updated_at" {
			t := v.(float64)
			r.UpdatedAt = time.Unix(int64(t), 0)
		} else if k == "id" {
			r.Id = v.(string)
		}
	}

	return r
}

// NewRecipeFromStreamRecord function handles converting a record from a
// DynamoDBStream Record into a Recipe
func NewTodoFromStreamRecord(record events.DynamoDBEventRecord) *Todo {
	r := &Todo{}
	for k, v := range record.Change.NewImage {
		if k == "id" {
			r.Id = v.String()
		} else if k == "description" {
			r.Description = v.String()
		} else if k == "name" {
			r.Name = v.String()
		} else if k == "created_at" {
			t := v.String()
			createdTime, err := time.Parse(time.RFC3339, t)
			if err != nil {
				logrus.Error("Error parsing CreatedTimestamp")
			} else {
				r.CreatedAt = createdTime
			}
		} else if k == "updated_at" {
			t := v.String()
			updatedTime, err := time.Parse(time.RFC3339, t)
			if err != nil {
				logrus.Error("Error parsing UpdatedTimestamp")
			} else {
				r.UpdatedAt = updatedTime
			}
		}
	}
	return r
}
