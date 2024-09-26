package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Codespace holds the schema definition for the Codespace entity.
type Codespace struct {
	ent.Schema
}

// Fields of the Codespace.
func (Codespace) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).StorageKey("old").Unique(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.String("machine_type").NotEmpty(),
		field.String("client_ip").NotEmpty(),
		field.String("github_user_id").NotEmpty(),
		field.String("github_codespace_id").NotEmpty(),
		field.String("github_codespace_url").NotEmpty(),
	}
}

// Edges of the Codespace.
func (Codespace) Edges() []ent.Edge {
	return nil
}
