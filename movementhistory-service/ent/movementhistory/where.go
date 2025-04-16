// Code generated by ent, DO NOT EDIT.

package movementhistory

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/Ostap00034/siproject-beercut-backend/movementhistory-service/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.MovementHistory {
	return predicate.MovementHistory(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.MovementHistory {
	return predicate.MovementHistory(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.MovementHistory {
	return predicate.MovementHistory(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.MovementHistory {
	return predicate.MovementHistory(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.MovementHistory {
	return predicate.MovementHistory(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.MovementHistory {
	return predicate.MovementHistory(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.MovementHistory {
	return predicate.MovementHistory(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.MovementHistory {
	return predicate.MovementHistory(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.MovementHistory {
	return predicate.MovementHistory(sql.FieldLTE(FieldID, id))
}

// PictureID applies equality check predicate on the "picture_id" field. It's identical to PictureIDEQ.
func PictureID(v string) predicate.MovementHistory {
	return predicate.MovementHistory(sql.FieldEQ(FieldPictureID, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v string) predicate.MovementHistory {
	return predicate.MovementHistory(sql.FieldEQ(FieldUserID, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.MovementHistory {
	return predicate.MovementHistory(sql.FieldEQ(FieldCreatedAt, v))
}

// PictureIDEQ applies the EQ predicate on the "picture_id" field.
func PictureIDEQ(v string) predicate.MovementHistory {
	return predicate.MovementHistory(sql.FieldEQ(FieldPictureID, v))
}

// PictureIDNEQ applies the NEQ predicate on the "picture_id" field.
func PictureIDNEQ(v string) predicate.MovementHistory {
	return predicate.MovementHistory(sql.FieldNEQ(FieldPictureID, v))
}

// PictureIDIn applies the In predicate on the "picture_id" field.
func PictureIDIn(vs ...string) predicate.MovementHistory {
	return predicate.MovementHistory(sql.FieldIn(FieldPictureID, vs...))
}

// PictureIDNotIn applies the NotIn predicate on the "picture_id" field.
func PictureIDNotIn(vs ...string) predicate.MovementHistory {
	return predicate.MovementHistory(sql.FieldNotIn(FieldPictureID, vs...))
}

// PictureIDGT applies the GT predicate on the "picture_id" field.
func PictureIDGT(v string) predicate.MovementHistory {
	return predicate.MovementHistory(sql.FieldGT(FieldPictureID, v))
}

// PictureIDGTE applies the GTE predicate on the "picture_id" field.
func PictureIDGTE(v string) predicate.MovementHistory {
	return predicate.MovementHistory(sql.FieldGTE(FieldPictureID, v))
}

// PictureIDLT applies the LT predicate on the "picture_id" field.
func PictureIDLT(v string) predicate.MovementHistory {
	return predicate.MovementHistory(sql.FieldLT(FieldPictureID, v))
}

// PictureIDLTE applies the LTE predicate on the "picture_id" field.
func PictureIDLTE(v string) predicate.MovementHistory {
	return predicate.MovementHistory(sql.FieldLTE(FieldPictureID, v))
}

// PictureIDContains applies the Contains predicate on the "picture_id" field.
func PictureIDContains(v string) predicate.MovementHistory {
	return predicate.MovementHistory(sql.FieldContains(FieldPictureID, v))
}

// PictureIDHasPrefix applies the HasPrefix predicate on the "picture_id" field.
func PictureIDHasPrefix(v string) predicate.MovementHistory {
	return predicate.MovementHistory(sql.FieldHasPrefix(FieldPictureID, v))
}

// PictureIDHasSuffix applies the HasSuffix predicate on the "picture_id" field.
func PictureIDHasSuffix(v string) predicate.MovementHistory {
	return predicate.MovementHistory(sql.FieldHasSuffix(FieldPictureID, v))
}

// PictureIDEqualFold applies the EqualFold predicate on the "picture_id" field.
func PictureIDEqualFold(v string) predicate.MovementHistory {
	return predicate.MovementHistory(sql.FieldEqualFold(FieldPictureID, v))
}

// PictureIDContainsFold applies the ContainsFold predicate on the "picture_id" field.
func PictureIDContainsFold(v string) predicate.MovementHistory {
	return predicate.MovementHistory(sql.FieldContainsFold(FieldPictureID, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v string) predicate.MovementHistory {
	return predicate.MovementHistory(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v string) predicate.MovementHistory {
	return predicate.MovementHistory(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...string) predicate.MovementHistory {
	return predicate.MovementHistory(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...string) predicate.MovementHistory {
	return predicate.MovementHistory(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v string) predicate.MovementHistory {
	return predicate.MovementHistory(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v string) predicate.MovementHistory {
	return predicate.MovementHistory(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v string) predicate.MovementHistory {
	return predicate.MovementHistory(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v string) predicate.MovementHistory {
	return predicate.MovementHistory(sql.FieldLTE(FieldUserID, v))
}

// UserIDContains applies the Contains predicate on the "user_id" field.
func UserIDContains(v string) predicate.MovementHistory {
	return predicate.MovementHistory(sql.FieldContains(FieldUserID, v))
}

// UserIDHasPrefix applies the HasPrefix predicate on the "user_id" field.
func UserIDHasPrefix(v string) predicate.MovementHistory {
	return predicate.MovementHistory(sql.FieldHasPrefix(FieldUserID, v))
}

// UserIDHasSuffix applies the HasSuffix predicate on the "user_id" field.
func UserIDHasSuffix(v string) predicate.MovementHistory {
	return predicate.MovementHistory(sql.FieldHasSuffix(FieldUserID, v))
}

// UserIDEqualFold applies the EqualFold predicate on the "user_id" field.
func UserIDEqualFold(v string) predicate.MovementHistory {
	return predicate.MovementHistory(sql.FieldEqualFold(FieldUserID, v))
}

// UserIDContainsFold applies the ContainsFold predicate on the "user_id" field.
func UserIDContainsFold(v string) predicate.MovementHistory {
	return predicate.MovementHistory(sql.FieldContainsFold(FieldUserID, v))
}

// FromEQ applies the EQ predicate on the "from" field.
func FromEQ(v From) predicate.MovementHistory {
	return predicate.MovementHistory(sql.FieldEQ(FieldFrom, v))
}

// FromNEQ applies the NEQ predicate on the "from" field.
func FromNEQ(v From) predicate.MovementHistory {
	return predicate.MovementHistory(sql.FieldNEQ(FieldFrom, v))
}

// FromIn applies the In predicate on the "from" field.
func FromIn(vs ...From) predicate.MovementHistory {
	return predicate.MovementHistory(sql.FieldIn(FieldFrom, vs...))
}

// FromNotIn applies the NotIn predicate on the "from" field.
func FromNotIn(vs ...From) predicate.MovementHistory {
	return predicate.MovementHistory(sql.FieldNotIn(FieldFrom, vs...))
}

// ToEQ applies the EQ predicate on the "to" field.
func ToEQ(v To) predicate.MovementHistory {
	return predicate.MovementHistory(sql.FieldEQ(FieldTo, v))
}

// ToNEQ applies the NEQ predicate on the "to" field.
func ToNEQ(v To) predicate.MovementHistory {
	return predicate.MovementHistory(sql.FieldNEQ(FieldTo, v))
}

// ToIn applies the In predicate on the "to" field.
func ToIn(vs ...To) predicate.MovementHistory {
	return predicate.MovementHistory(sql.FieldIn(FieldTo, vs...))
}

// ToNotIn applies the NotIn predicate on the "to" field.
func ToNotIn(vs ...To) predicate.MovementHistory {
	return predicate.MovementHistory(sql.FieldNotIn(FieldTo, vs...))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.MovementHistory {
	return predicate.MovementHistory(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.MovementHistory {
	return predicate.MovementHistory(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.MovementHistory {
	return predicate.MovementHistory(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.MovementHistory {
	return predicate.MovementHistory(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.MovementHistory {
	return predicate.MovementHistory(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.MovementHistory {
	return predicate.MovementHistory(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.MovementHistory {
	return predicate.MovementHistory(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.MovementHistory {
	return predicate.MovementHistory(sql.FieldLTE(FieldCreatedAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.MovementHistory) predicate.MovementHistory {
	return predicate.MovementHistory(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.MovementHistory) predicate.MovementHistory {
	return predicate.MovementHistory(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.MovementHistory) predicate.MovementHistory {
	return predicate.MovementHistory(sql.NotPredicates(p))
}
