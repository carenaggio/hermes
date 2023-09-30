/*
Copyright 2023 Christos Triantafyllidis <christos.triantafyllidis@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// System holds the schema definition for the System entity.
type System struct {
	ent.Schema
}

// Fields of the System.
func (System) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("system_id", uuid.New()).Unique(),
		field.String("public_key"),
		field.Bool("approved"),
		field.Int64("last_login"),
	}
}

// Edges of the System.
func (System) Edges() []ent.Edge {
	return nil
}
