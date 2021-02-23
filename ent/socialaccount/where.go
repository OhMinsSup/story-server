// Code generated by entc, DO NOT EDIT.

package socialaccount

import (
	"time"

	"github.com/OhMinsSup/story-server/ent/predicate"
	"github.com/facebook/ent/dialect/sql"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// SocialID applies equality check predicate on the "social_id" field. It's identical to SocialIDEQ.
func SocialID(v string) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSocialID), v))
	})
}

// AccessToken applies equality check predicate on the "access_token" field. It's identical to AccessTokenEQ.
func AccessToken(v string) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAccessToken), v))
	})
}

// Provider applies equality check predicate on the "provider" field. It's identical to ProviderEQ.
func Provider(v string) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProvider), v))
	})
}

// FkUserID applies equality check predicate on the "fk_user_id" field. It's identical to FkUserIDEQ.
func FkUserID(v uuid.UUID) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFkUserID), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// SocialIDEQ applies the EQ predicate on the "social_id" field.
func SocialIDEQ(v string) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSocialID), v))
	})
}

// SocialIDNEQ applies the NEQ predicate on the "social_id" field.
func SocialIDNEQ(v string) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSocialID), v))
	})
}

// SocialIDIn applies the In predicate on the "social_id" field.
func SocialIDIn(vs ...string) predicate.SocialAccount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SocialAccount(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSocialID), v...))
	})
}

// SocialIDNotIn applies the NotIn predicate on the "social_id" field.
func SocialIDNotIn(vs ...string) predicate.SocialAccount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SocialAccount(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSocialID), v...))
	})
}

// SocialIDGT applies the GT predicate on the "social_id" field.
func SocialIDGT(v string) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSocialID), v))
	})
}

// SocialIDGTE applies the GTE predicate on the "social_id" field.
func SocialIDGTE(v string) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSocialID), v))
	})
}

// SocialIDLT applies the LT predicate on the "social_id" field.
func SocialIDLT(v string) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSocialID), v))
	})
}

// SocialIDLTE applies the LTE predicate on the "social_id" field.
func SocialIDLTE(v string) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSocialID), v))
	})
}

// SocialIDContains applies the Contains predicate on the "social_id" field.
func SocialIDContains(v string) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSocialID), v))
	})
}

// SocialIDHasPrefix applies the HasPrefix predicate on the "social_id" field.
func SocialIDHasPrefix(v string) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSocialID), v))
	})
}

// SocialIDHasSuffix applies the HasSuffix predicate on the "social_id" field.
func SocialIDHasSuffix(v string) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSocialID), v))
	})
}

// SocialIDEqualFold applies the EqualFold predicate on the "social_id" field.
func SocialIDEqualFold(v string) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSocialID), v))
	})
}

// SocialIDContainsFold applies the ContainsFold predicate on the "social_id" field.
func SocialIDContainsFold(v string) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSocialID), v))
	})
}

// AccessTokenEQ applies the EQ predicate on the "access_token" field.
func AccessTokenEQ(v string) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAccessToken), v))
	})
}

// AccessTokenNEQ applies the NEQ predicate on the "access_token" field.
func AccessTokenNEQ(v string) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAccessToken), v))
	})
}

// AccessTokenIn applies the In predicate on the "access_token" field.
func AccessTokenIn(vs ...string) predicate.SocialAccount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SocialAccount(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAccessToken), v...))
	})
}

// AccessTokenNotIn applies the NotIn predicate on the "access_token" field.
func AccessTokenNotIn(vs ...string) predicate.SocialAccount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SocialAccount(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAccessToken), v...))
	})
}

// AccessTokenGT applies the GT predicate on the "access_token" field.
func AccessTokenGT(v string) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAccessToken), v))
	})
}

// AccessTokenGTE applies the GTE predicate on the "access_token" field.
func AccessTokenGTE(v string) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAccessToken), v))
	})
}

// AccessTokenLT applies the LT predicate on the "access_token" field.
func AccessTokenLT(v string) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAccessToken), v))
	})
}

// AccessTokenLTE applies the LTE predicate on the "access_token" field.
func AccessTokenLTE(v string) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAccessToken), v))
	})
}

// AccessTokenContains applies the Contains predicate on the "access_token" field.
func AccessTokenContains(v string) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAccessToken), v))
	})
}

// AccessTokenHasPrefix applies the HasPrefix predicate on the "access_token" field.
func AccessTokenHasPrefix(v string) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAccessToken), v))
	})
}

// AccessTokenHasSuffix applies the HasSuffix predicate on the "access_token" field.
func AccessTokenHasSuffix(v string) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAccessToken), v))
	})
}

// AccessTokenEqualFold applies the EqualFold predicate on the "access_token" field.
func AccessTokenEqualFold(v string) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAccessToken), v))
	})
}

// AccessTokenContainsFold applies the ContainsFold predicate on the "access_token" field.
func AccessTokenContainsFold(v string) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAccessToken), v))
	})
}

// ProviderEQ applies the EQ predicate on the "provider" field.
func ProviderEQ(v string) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProvider), v))
	})
}

// ProviderNEQ applies the NEQ predicate on the "provider" field.
func ProviderNEQ(v string) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldProvider), v))
	})
}

// ProviderIn applies the In predicate on the "provider" field.
func ProviderIn(vs ...string) predicate.SocialAccount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SocialAccount(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldProvider), v...))
	})
}

// ProviderNotIn applies the NotIn predicate on the "provider" field.
func ProviderNotIn(vs ...string) predicate.SocialAccount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SocialAccount(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldProvider), v...))
	})
}

// ProviderGT applies the GT predicate on the "provider" field.
func ProviderGT(v string) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldProvider), v))
	})
}

// ProviderGTE applies the GTE predicate on the "provider" field.
func ProviderGTE(v string) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldProvider), v))
	})
}

// ProviderLT applies the LT predicate on the "provider" field.
func ProviderLT(v string) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldProvider), v))
	})
}

// ProviderLTE applies the LTE predicate on the "provider" field.
func ProviderLTE(v string) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldProvider), v))
	})
}

// ProviderContains applies the Contains predicate on the "provider" field.
func ProviderContains(v string) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldProvider), v))
	})
}

// ProviderHasPrefix applies the HasPrefix predicate on the "provider" field.
func ProviderHasPrefix(v string) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldProvider), v))
	})
}

// ProviderHasSuffix applies the HasSuffix predicate on the "provider" field.
func ProviderHasSuffix(v string) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldProvider), v))
	})
}

// ProviderEqualFold applies the EqualFold predicate on the "provider" field.
func ProviderEqualFold(v string) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldProvider), v))
	})
}

// ProviderContainsFold applies the ContainsFold predicate on the "provider" field.
func ProviderContainsFold(v string) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldProvider), v))
	})
}

// FkUserIDEQ applies the EQ predicate on the "fk_user_id" field.
func FkUserIDEQ(v uuid.UUID) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFkUserID), v))
	})
}

// FkUserIDNEQ applies the NEQ predicate on the "fk_user_id" field.
func FkUserIDNEQ(v uuid.UUID) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFkUserID), v))
	})
}

// FkUserIDIn applies the In predicate on the "fk_user_id" field.
func FkUserIDIn(vs ...uuid.UUID) predicate.SocialAccount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SocialAccount(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFkUserID), v...))
	})
}

// FkUserIDNotIn applies the NotIn predicate on the "fk_user_id" field.
func FkUserIDNotIn(vs ...uuid.UUID) predicate.SocialAccount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SocialAccount(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFkUserID), v...))
	})
}

// FkUserIDGT applies the GT predicate on the "fk_user_id" field.
func FkUserIDGT(v uuid.UUID) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFkUserID), v))
	})
}

// FkUserIDGTE applies the GTE predicate on the "fk_user_id" field.
func FkUserIDGTE(v uuid.UUID) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFkUserID), v))
	})
}

// FkUserIDLT applies the LT predicate on the "fk_user_id" field.
func FkUserIDLT(v uuid.UUID) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFkUserID), v))
	})
}

// FkUserIDLTE applies the LTE predicate on the "fk_user_id" field.
func FkUserIDLTE(v uuid.UUID) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFkUserID), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.SocialAccount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SocialAccount(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.SocialAccount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SocialAccount(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.SocialAccount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SocialAccount(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.SocialAccount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SocialAccount(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SocialAccount) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SocialAccount) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.SocialAccount) predicate.SocialAccount {
	return predicate.SocialAccount(func(s *sql.Selector) {
		p(s.Not())
	})
}
