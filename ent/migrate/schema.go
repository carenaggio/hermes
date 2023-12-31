// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// SystemsColumns holds the columns for the "systems" table.
	SystemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "system_id", Type: field.TypeUUID, Unique: true},
		{Name: "public_key", Type: field.TypeBytes},
		{Name: "approved", Type: field.TypeBool},
		{Name: "last_login", Type: field.TypeInt64},
	}
	// SystemsTable holds the schema information for the "systems" table.
	SystemsTable = &schema.Table{
		Name:       "systems",
		Columns:    SystemsColumns,
		PrimaryKey: []*schema.Column{SystemsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		SystemsTable,
	}
)

func init() {
}
