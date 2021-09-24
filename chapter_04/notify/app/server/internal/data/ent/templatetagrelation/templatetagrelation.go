// Code generated by entc, DO NOT EDIT.

package templatetagrelation

const (
	// Label holds the string label denoting the templatetagrelation type in the database.
	Label = "template_tag_relation"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// Table holds the table name of the templatetagrelation in the database.
	Table = "template_tag_relations"
)

// Columns holds all SQL columns for templatetagrelation fields.
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