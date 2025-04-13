// Code generated by ent, DO NOT EDIT.

package author

import (
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the author type in the database.
	Label = "author"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldFullName holds the string denoting the full_name field in the database.
	FieldFullName = "full_name"
	// FieldDateOfBirth holds the string denoting the date_of_birth field in the database.
	FieldDateOfBirth = "date_of_birth"
	// FieldDateOfDeath holds the string denoting the date_of_death field in the database.
	FieldDateOfDeath = "date_of_death"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// Table holds the table name of the author in the database.
	Table = "authors"
)

// Columns holds all SQL columns for author fields.
var Columns = []string{
	FieldID,
	FieldFullName,
	FieldDateOfBirth,
	FieldDateOfDeath,
	FieldCreatedAt,
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
	DefaultCreatedAt time.Time
)

// OrderOption defines the ordering options for the Author queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByFullName orders the results by the full_name field.
func ByFullName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFullName, opts...).ToFunc()
}

// ByDateOfBirth orders the results by the date_of_birth field.
func ByDateOfBirth(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDateOfBirth, opts...).ToFunc()
}

// ByDateOfDeath orders the results by the date_of_death field.
func ByDateOfDeath(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDateOfDeath, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}
