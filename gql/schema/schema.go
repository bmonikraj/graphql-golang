package schema

import (
	"log"
	"github.com/graphql-go/graphql"
	"github.com/bmonikraj/goql/gql/types"
)

func InitSchema() graphql.Schema {
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: types.FieldType}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery), Mutation: types.MutationType}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error : %v", err)
	}
	return schema
}