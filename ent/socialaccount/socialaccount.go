// Code generated by entc, DO NOT EDIT.

package socialaccount

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the socialaccount type in the database.
	Label = "social_account"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSocialID holds the string denoting the social_id field in the database.
	FieldSocialID = "social_id"
	// FieldAccessToken holds the string denoting the access_token field in the database.
	FieldAccessToken = "access_token"
	// FieldProvider holds the string denoting the provider field in the database.
	FieldProvider = "provider"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"

	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"

	// Table holds the table name of the socialaccount in the database.
	Table = "social_accounts"
	// UserTable is the table the holds the user relation/edge.
	UserTable = "social_accounts"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "fk_user_id"
)

// Columns holds all SQL columns for socialaccount fields.
var Columns = []string{
	FieldID,
	FieldSocialID,
	FieldAccessToken,
	FieldProvider,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the SocialAccount type.
var ForeignKeys = []string{
	"fk_user_id",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// SocialIDValidator is a validator for the "social_id" field. It is called by the builders before save.
	SocialIDValidator func(string) error
	// AccessTokenValidator is a validator for the "access_token" field. It is called by the builders before save.
	AccessTokenValidator func(string) error
	// ProviderValidator is a validator for the "provider" field. It is called by the builders before save.
	ProviderValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)