// Code generated by ent, DO NOT EDIT.

package offer

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the offer type in the database.
	Label = "offer"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldMasterID holds the string denoting the master_id field in the database.
	FieldMasterID = "master_id"
	// FieldOrderID holds the string denoting the order_id field in the database.
	FieldOrderID = "order_id"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldPrice holds the string denoting the price field in the database.
	FieldPrice = "price"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// Table holds the table name of the offer in the database.
	Table = "offers"
)

// Columns holds all SQL columns for offer fields.
var Columns = []string{
	FieldID,
	FieldMasterID,
	FieldOrderID,
	FieldStatus,
	FieldPrice,
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
	// DefaultPrice holds the default value on creation for the "price" field.
	DefaultPrice float32
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Status defines the type for the "status" enum field.
type Status string

// StatusActive is the default value of the Status enum.
const DefaultStatus = StatusActive

// Status values.
const (
	StatusActive    Status = "active"
	StatusAccepted  Status = "accepted"
	StatusCancelled Status = "cancelled"
	StatusRejected  Status = "rejected"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusActive, StatusAccepted, StatusCancelled, StatusRejected:
		return nil
	default:
		return fmt.Errorf("offer: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the Offer queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByMasterID orders the results by the master_id field.
func ByMasterID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMasterID, opts...).ToFunc()
}

// ByOrderID orders the results by the order_id field.
func ByOrderID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOrderID, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByPrice orders the results by the price field.
func ByPrice(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPrice, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}
