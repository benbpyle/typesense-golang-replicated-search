package lib

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/typesense/typesense-go/v2/typesense"
	"github.com/typesense/typesense-go/v2/typesense/api"
)

func SearchDocuments(ctx context.Context, client *typesense.Client, query string) ([]Recipe, error) {
	queryBy := "name"
	sortBy := "createdTimestamp:desc"
	searchParameters := &api.SearchCollectionParams{
		Q:       &query,
		QueryBy: &queryBy,
		SortBy:  &sortBy,
	}

	results, err := client.Collection("recipes").Documents().Search(ctx, searchParameters)
	if err != nil {
		return nil, err
	}

	recipes := []Recipe{}
	for _, v := range *results.Hits {
		logrus.Infof("Docs: %v", v.Document)
		r := NewRecipeFromTypesenseRecipe(*v.Document)

		recipes = append(recipes, *r)
	}

	return recipes, nil
}
