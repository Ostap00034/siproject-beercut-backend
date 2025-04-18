// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/Ostap00034/siproject-beercut-backend/picture-service/ent/picture"
)

// Picture is the model entity for the Picture schema.
type Picture struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// DateOfPainting holds the value of the "date_of_painting" field.
	DateOfPainting string `json:"date_of_painting,omitempty"`
	// GenresIds holds the value of the "genres_ids" field.
	GenresIds []string `json:"genres_ids,omitempty"`
	// AuthorsIds holds the value of the "authors_ids" field.
	AuthorsIds []string `json:"authors_ids,omitempty"`
	// ExhibitionID holds the value of the "exhibition_id" field.
	ExhibitionID string `json:"exhibition_id,omitempty"`
	// Cost holds the value of the "cost" field.
	Cost float64 `json:"cost,omitempty"`
	// Location holds the value of the "location" field.
	Location picture.Location `json:"location,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt    time.Time `json:"created_at,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Picture) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case picture.FieldGenresIds, picture.FieldAuthorsIds:
			values[i] = new([]byte)
		case picture.FieldCost:
			values[i] = new(sql.NullFloat64)
		case picture.FieldID:
			values[i] = new(sql.NullInt64)
		case picture.FieldName, picture.FieldDateOfPainting, picture.FieldExhibitionID, picture.FieldLocation:
			values[i] = new(sql.NullString)
		case picture.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Picture fields.
func (pi *Picture) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case picture.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pi.ID = int(value.Int64)
		case picture.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pi.Name = value.String
			}
		case picture.FieldDateOfPainting:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field date_of_painting", values[i])
			} else if value.Valid {
				pi.DateOfPainting = value.String
			}
		case picture.FieldGenresIds:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field genres_ids", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pi.GenresIds); err != nil {
					return fmt.Errorf("unmarshal field genres_ids: %w", err)
				}
			}
		case picture.FieldAuthorsIds:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field authors_ids", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pi.AuthorsIds); err != nil {
					return fmt.Errorf("unmarshal field authors_ids: %w", err)
				}
			}
		case picture.FieldExhibitionID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field exhibition_id", values[i])
			} else if value.Valid {
				pi.ExhibitionID = value.String
			}
		case picture.FieldCost:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field cost", values[i])
			} else if value.Valid {
				pi.Cost = value.Float64
			}
		case picture.FieldLocation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field location", values[i])
			} else if value.Valid {
				pi.Location = picture.Location(value.String)
			}
		case picture.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pi.CreatedAt = value.Time
			}
		default:
			pi.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Picture.
// This includes values selected through modifiers, order, etc.
func (pi *Picture) Value(name string) (ent.Value, error) {
	return pi.selectValues.Get(name)
}

// Update returns a builder for updating this Picture.
// Note that you need to call Picture.Unwrap() before calling this method if this Picture
// was returned from a transaction, and the transaction was committed or rolled back.
func (pi *Picture) Update() *PictureUpdateOne {
	return NewPictureClient(pi.config).UpdateOne(pi)
}

// Unwrap unwraps the Picture entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pi *Picture) Unwrap() *Picture {
	_tx, ok := pi.config.driver.(*txDriver)
	if !ok {
		panic("ent: Picture is not a transactional entity")
	}
	pi.config.driver = _tx.drv
	return pi
}

// String implements the fmt.Stringer.
func (pi *Picture) String() string {
	var builder strings.Builder
	builder.WriteString("Picture(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pi.ID))
	builder.WriteString("name=")
	builder.WriteString(pi.Name)
	builder.WriteString(", ")
	builder.WriteString("date_of_painting=")
	builder.WriteString(pi.DateOfPainting)
	builder.WriteString(", ")
	builder.WriteString("genres_ids=")
	builder.WriteString(fmt.Sprintf("%v", pi.GenresIds))
	builder.WriteString(", ")
	builder.WriteString("authors_ids=")
	builder.WriteString(fmt.Sprintf("%v", pi.AuthorsIds))
	builder.WriteString(", ")
	builder.WriteString("exhibition_id=")
	builder.WriteString(pi.ExhibitionID)
	builder.WriteString(", ")
	builder.WriteString("cost=")
	builder.WriteString(fmt.Sprintf("%v", pi.Cost))
	builder.WriteString(", ")
	builder.WriteString("location=")
	builder.WriteString(fmt.Sprintf("%v", pi.Location))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(pi.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Pictures is a parsable slice of Picture.
type Pictures []*Picture
