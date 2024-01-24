package enthistory

import (
	"errors"
	"regexp"
	"strings"

	"entgo.io/ent/schema/field"

	"entgo.io/ent/entc/load"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func toSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func copyRef[T any](ref *T) *T {
	if ref == nil {
		return nil
	}
	val := *ref
	return &val
}

func loadHistorySchema(IdType string) (*load.Schema, error) {
	schema := history{}

	switch IdType {
	case "int":
		schema.ref = field.Int("ref").Immutable().Optional()
	case "string":
		schema.ref = field.String("ref").Immutable().Optional()
	default:
		return nil, errors.New("only id and string are supported id types right now")
	}

	bytes, err := load.MarshalSchema(schema)
	if err != nil {
		return nil, err
	}

	historySchema, err := load.UnmarshalSchema(bytes)
	if err != nil {
		return nil, err
	}

	return historySchema, nil
}

func getUpdatedByField(updatedByValueType string) (*load.Field, error) {
	if updatedByValueType == "String" {
		return load.NewField(field.String("updated_by").Optional().Nillable().Immutable().Descriptor())
	}
	if updatedByValueType == "Int" {
		return load.NewField(field.Int("updated_by").Optional().Nillable().Immutable().Descriptor())
	}
	return nil, nil
}

func getDeletedByField(deletedByValueType string) (*load.Field, error) {
	if deletedByValueType == "String" {
		return load.NewField(field.String("deleted_by").Optional().Nillable().Immutable().Descriptor())
	}
	if deletedByValueType == "Int" {
		return load.NewField(field.Int("deleted_by").Optional().Nillable().Immutable().Descriptor())
	}
	return nil, nil
}

func getHistoryAnnotations(schema *load.Schema) Annotations {
	annotations := Annotations{}
	if historyAnnotations, ok := schema.Annotations["History"].(map[string]any); ok {
		if exclude, ok := historyAnnotations["exclude"].(bool); ok {
			annotations.Exclude = exclude
		}
		if isHistory, ok := historyAnnotations["isHistory"].(bool); ok {
			annotations.IsHistory = isHistory
		}
	}
	return annotations
}

func getSchemaTableName(schema *load.Schema) string {
	if entSqlMap, ok := schema.Annotations["EntSQL"].(map[string]any); ok {
		if table, ok := entSqlMap["table"].(string); ok {
			return table
		}
	}
	return toSnakeCase(schema.Name)
}

func getPkgFromSchemaPath(schemaPath string) (string, error) {
	parts := strings.Split(schemaPath, "/")
	lastPart := parts[len(parts)-1]
	if len(lastPart) == 0 {
		return "", errors.New("invalid schema path, unable to find package name in path")
	}
	return lastPart, nil
}
