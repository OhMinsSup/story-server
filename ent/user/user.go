// Code generated by entc, DO NOT EDIT.

package user

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldIsCertified holds the string denoting the is_certified field in the database.
	FieldIsCertified = "is_certified"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"

	// EdgeUserProfile holds the string denoting the user_profile edge name in mutations.
	EdgeUserProfile = "user_profile"
	// EdgeVelogConfig holds the string denoting the velog_config edge name in mutations.
	EdgeVelogConfig = "velog_config"
	// EdgeUserMeta holds the string denoting the user_meta edge name in mutations.
	EdgeUserMeta = "user_meta"

	// Table holds the table name of the user in the database.
	Table = "users"
	// UserProfileTable is the table the holds the user_profile relation/edge.
	UserProfileTable = "user_profiles"
	// UserProfileInverseTable is the table name for the UserProfile entity.
	// It exists in this package in order to avoid circular dependency with the "userprofile" package.
	UserProfileInverseTable = "user_profiles"
	// UserProfileColumn is the table column denoting the user_profile relation/edge.
	UserProfileColumn = "fk_user_id"
	// VelogConfigTable is the table the holds the velog_config relation/edge.
	VelogConfigTable = "velog_configs"
	// VelogConfigInverseTable is the table name for the VelogConfig entity.
	// It exists in this package in order to avoid circular dependency with the "velogconfig" package.
	VelogConfigInverseTable = "velog_configs"
	// VelogConfigColumn is the table column denoting the velog_config relation/edge.
	VelogConfigColumn = "fk_user_id"
	// UserMetaTable is the table the holds the user_meta relation/edge.
	UserMetaTable = "user_meta"
	// UserMetaInverseTable is the table name for the UserMeta entity.
	// It exists in this package in order to avoid circular dependency with the "usermeta" package.
	UserMetaInverseTable = "user_meta"
	// UserMetaColumn is the table column denoting the user_meta relation/edge.
	UserMetaColumn = "fk_user_id"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldUsername,
	FieldEmail,
	FieldIsCertified,
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
	// UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	UsernameValidator func(string) error
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// DefaultIsCertified holds the default value on creation for the "is_certified" field.
	DefaultIsCertified bool
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
