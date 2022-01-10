// Code generated by entc, DO NOT EDIT.

package lilyskill

import (
	"time"
)

const (
	// Label holds the string label denoting the lilyskill type in the database.
	Label = "lily_skill"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldLilyID holds the string denoting the lily_id field in the database.
	FieldLilyID = "lily_id"
	// FieldSkillID holds the string denoting the skill_id field in the database.
	FieldSkillID = "skill_id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// Table holds the table name of the lilyskill in the database.
	Table = "lily_skills"
)

// Columns holds all SQL columns for lilyskill fields.
var Columns = []string{
	FieldID,
	FieldLilyID,
	FieldSkillID,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)