// Code generated by entc, DO NOT EDIT.

package templatetagrelation

import (
	"time"
)

const (
	// Label holds the string label denoting the templatetagrelation type in the database.
	Label = "template_tag_relation"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldTemplateUUID holds the string denoting the template_uuid field in the database.
	FieldTemplateUUID = "template_uuid"
	// FieldTagUUID holds the string denoting the tag_uuid field in the database.
	FieldTagUUID = "tag_uuid"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// Table holds the table name of the templatetagrelation in the database.
	Table = "notify_tag_template"
)

// Columns holds all SQL columns for templatetagrelation fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldTemplateUUID,
	FieldTagUUID,
	FieldStatus,
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
	// TemplateUUIDValidator is a validator for the "template_uuid" field. It is called by the builders before save.
	TemplateUUIDValidator func(string) error
	// TagUUIDValidator is a validator for the "tag_uuid" field. It is called by the builders before save.
	TagUUIDValidator func(string) error
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus int
	// StatusValidator is a validator for the "status" field. It is called by the builders before save.
	StatusValidator func(int) error
)
