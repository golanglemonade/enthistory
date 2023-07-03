// Code generated by enthistory, DO NOT EDIT.
package otherschema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"

	"github.com/flume/enthistory"

	"time"
)

// FriendshipHistory holds the schema definition for the FriendshipHistory entity.
type FriendshipHistory struct {
	ent.Schema
}

// Annotations of the FriendshipHistory.
func (FriendshipHistory) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "friendship_history",
		},
		enthistory.Annotations{
			IsHistory: true,
			Exclude:   true,
		},
	}
}

// Fields of the FriendshipHistory.
func (FriendshipHistory) Fields() []ent.Field {
	historyFields := []ent.Field{
		field.Time("history_time").
			Default(time.Now).
			Immutable(),
		field.Int("ref").
			Immutable().
			Optional(),
		field.Enum("operation").
			GoType(enthistory.OpType("")).
			Immutable(),
	}

	original := Friendship{}
	for _, field := range original.Fields() {
		if field.Descriptor().Name != "id" {
			historyFields = append(historyFields, field)
		}
	}

	return historyFields
}

// Mixin of the FriendshipHistory.
func (FriendshipHistory) Mixin() []ent.Mixin {
	return Friendship{}.Mixin()
}
