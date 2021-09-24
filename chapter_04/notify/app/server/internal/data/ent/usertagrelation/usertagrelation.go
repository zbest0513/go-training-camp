// Code generated by entc, DO NOT EDIT.

package usertagrelation

const (
	// Label holds the string label denoting the usertagrelation type in the database.
	Label = "user_tag_relation"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// Table holds the table name of the usertagrelation in the database.
	Table = "user_tag_relations"
)

// Columns holds all SQL columns for usertagrelation fields.
var Columns = []string{
	FieldID,
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