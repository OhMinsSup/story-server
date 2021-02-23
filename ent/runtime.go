// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/OhMinsSup/story-server/ent/authtoken"
	"github.com/OhMinsSup/story-server/ent/emailauth"
	"github.com/OhMinsSup/story-server/ent/schema"
	"github.com/OhMinsSup/story-server/ent/socialaccount"
	"github.com/OhMinsSup/story-server/ent/user"
	"github.com/OhMinsSup/story-server/ent/usermeta"
	"github.com/OhMinsSup/story-server/ent/userprofile"
	"github.com/OhMinsSup/story-server/ent/velogconfig"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	authtokenFields := schema.AuthToken{}.Fields()
	_ = authtokenFields
	// authtokenDescDisabled is the schema descriptor for disabled field.
	authtokenDescDisabled := authtokenFields[1].Descriptor()
	// authtoken.DefaultDisabled holds the default value on creation for the disabled field.
	authtoken.DefaultDisabled = authtokenDescDisabled.Default.(bool)
	// authtokenDescCreatedAt is the schema descriptor for created_at field.
	authtokenDescCreatedAt := authtokenFields[2].Descriptor()
	// authtoken.DefaultCreatedAt holds the default value on creation for the created_at field.
	authtoken.DefaultCreatedAt = authtokenDescCreatedAt.Default.(func() time.Time)
	// authtokenDescUpdatedAt is the schema descriptor for updated_at field.
	authtokenDescUpdatedAt := authtokenFields[3].Descriptor()
	// authtoken.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	authtoken.DefaultUpdatedAt = authtokenDescUpdatedAt.Default.(func() time.Time)
	// authtoken.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	authtoken.UpdateDefaultUpdatedAt = authtokenDescUpdatedAt.UpdateDefault.(func() time.Time)
	// authtokenDescID is the schema descriptor for id field.
	authtokenDescID := authtokenFields[0].Descriptor()
	// authtoken.DefaultID holds the default value on creation for the id field.
	authtoken.DefaultID = authtokenDescID.Default.(func() uuid.UUID)
	emailauthFields := schema.EmailAuth{}.Fields()
	_ = emailauthFields
	// emailauthDescLogged is the schema descriptor for logged field.
	emailauthDescLogged := emailauthFields[3].Descriptor()
	// emailauth.DefaultLogged holds the default value on creation for the logged field.
	emailauth.DefaultLogged = emailauthDescLogged.Default.(bool)
	// emailauthDescCreatedAt is the schema descriptor for created_at field.
	emailauthDescCreatedAt := emailauthFields[4].Descriptor()
	// emailauth.DefaultCreatedAt holds the default value on creation for the created_at field.
	emailauth.DefaultCreatedAt = emailauthDescCreatedAt.Default.(func() time.Time)
	// emailauthDescUpdatedAt is the schema descriptor for updated_at field.
	emailauthDescUpdatedAt := emailauthFields[5].Descriptor()
	// emailauth.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	emailauth.DefaultUpdatedAt = emailauthDescUpdatedAt.Default.(func() time.Time)
	// emailauth.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	emailauth.UpdateDefaultUpdatedAt = emailauthDescUpdatedAt.UpdateDefault.(func() time.Time)
	// emailauthDescID is the schema descriptor for id field.
	emailauthDescID := emailauthFields[0].Descriptor()
	// emailauth.DefaultID holds the default value on creation for the id field.
	emailauth.DefaultID = emailauthDescID.Default.(func() uuid.UUID)
	socialaccountFields := schema.SocialAccount{}.Fields()
	_ = socialaccountFields
	// socialaccountDescSocialID is the schema descriptor for social_id field.
	socialaccountDescSocialID := socialaccountFields[1].Descriptor()
	// socialaccount.SocialIDValidator is a validator for the "social_id" field. It is called by the builders before save.
	socialaccount.SocialIDValidator = socialaccountDescSocialID.Validators[0].(func(string) error)
	// socialaccountDescAccessToken is the schema descriptor for access_token field.
	socialaccountDescAccessToken := socialaccountFields[2].Descriptor()
	// socialaccount.AccessTokenValidator is a validator for the "access_token" field. It is called by the builders before save.
	socialaccount.AccessTokenValidator = socialaccountDescAccessToken.Validators[0].(func(string) error)
	// socialaccountDescProvider is the schema descriptor for provider field.
	socialaccountDescProvider := socialaccountFields[3].Descriptor()
	// socialaccount.ProviderValidator is a validator for the "provider" field. It is called by the builders before save.
	socialaccount.ProviderValidator = socialaccountDescProvider.Validators[0].(func(string) error)
	// socialaccountDescCreatedAt is the schema descriptor for created_at field.
	socialaccountDescCreatedAt := socialaccountFields[5].Descriptor()
	// socialaccount.DefaultCreatedAt holds the default value on creation for the created_at field.
	socialaccount.DefaultCreatedAt = socialaccountDescCreatedAt.Default.(func() time.Time)
	// socialaccountDescUpdatedAt is the schema descriptor for updated_at field.
	socialaccountDescUpdatedAt := socialaccountFields[6].Descriptor()
	// socialaccount.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	socialaccount.DefaultUpdatedAt = socialaccountDescUpdatedAt.Default.(func() time.Time)
	// socialaccount.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	socialaccount.UpdateDefaultUpdatedAt = socialaccountDescUpdatedAt.UpdateDefault.(func() time.Time)
	// socialaccountDescID is the schema descriptor for id field.
	socialaccountDescID := socialaccountFields[0].Descriptor()
	// socialaccount.DefaultID holds the default value on creation for the id field.
	socialaccount.DefaultID = socialaccountDescID.Default.(func() uuid.UUID)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[1].Descriptor()
	// user.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	user.UsernameValidator = userDescUsername.Validators[0].(func(string) error)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[2].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescIsCertified is the schema descriptor for is_certified field.
	userDescIsCertified := userFields[3].Descriptor()
	// user.DefaultIsCertified holds the default value on creation for the is_certified field.
	user.DefaultIsCertified = userDescIsCertified.Default.(bool)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[4].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[5].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
	usermetaFields := schema.UserMeta{}.Fields()
	_ = usermetaFields
	// usermetaDescEmailNotification is the schema descriptor for email_notification field.
	usermetaDescEmailNotification := usermetaFields[1].Descriptor()
	// usermeta.DefaultEmailNotification holds the default value on creation for the email_notification field.
	usermeta.DefaultEmailNotification = usermetaDescEmailNotification.Default.(bool)
	// usermetaDescEmailPromotions is the schema descriptor for email_promotions field.
	usermetaDescEmailPromotions := usermetaFields[2].Descriptor()
	// usermeta.DefaultEmailPromotions holds the default value on creation for the email_promotions field.
	usermeta.DefaultEmailPromotions = usermetaDescEmailPromotions.Default.(bool)
	// usermetaDescCreatedAt is the schema descriptor for created_at field.
	usermetaDescCreatedAt := usermetaFields[3].Descriptor()
	// usermeta.DefaultCreatedAt holds the default value on creation for the created_at field.
	usermeta.DefaultCreatedAt = usermetaDescCreatedAt.Default.(func() time.Time)
	// usermetaDescUpdatedAt is the schema descriptor for updated_at field.
	usermetaDescUpdatedAt := usermetaFields[4].Descriptor()
	// usermeta.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	usermeta.DefaultUpdatedAt = usermetaDescUpdatedAt.Default.(func() time.Time)
	// usermeta.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	usermeta.UpdateDefaultUpdatedAt = usermetaDescUpdatedAt.UpdateDefault.(func() time.Time)
	// usermetaDescID is the schema descriptor for id field.
	usermetaDescID := usermetaFields[0].Descriptor()
	// usermeta.DefaultID holds the default value on creation for the id field.
	usermeta.DefaultID = usermetaDescID.Default.(func() uuid.UUID)
	userprofileFields := schema.UserProfile{}.Fields()
	_ = userprofileFields
	// userprofileDescDisplayName is the schema descriptor for display_name field.
	userprofileDescDisplayName := userprofileFields[1].Descriptor()
	// userprofile.DisplayNameValidator is a validator for the "display_name" field. It is called by the builders before save.
	userprofile.DisplayNameValidator = userprofileDescDisplayName.Validators[0].(func(string) error)
	// userprofileDescShortBio is the schema descriptor for short_bio field.
	userprofileDescShortBio := userprofileFields[2].Descriptor()
	// userprofile.ShortBioValidator is a validator for the "short_bio" field. It is called by the builders before save.
	userprofile.ShortBioValidator = userprofileDescShortBio.Validators[0].(func(string) error)
	// userprofileDescCreatedAt is the schema descriptor for created_at field.
	userprofileDescCreatedAt := userprofileFields[6].Descriptor()
	// userprofile.DefaultCreatedAt holds the default value on creation for the created_at field.
	userprofile.DefaultCreatedAt = userprofileDescCreatedAt.Default.(func() time.Time)
	// userprofileDescUpdatedAt is the schema descriptor for updated_at field.
	userprofileDescUpdatedAt := userprofileFields[7].Descriptor()
	// userprofile.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	userprofile.DefaultUpdatedAt = userprofileDescUpdatedAt.Default.(func() time.Time)
	// userprofile.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	userprofile.UpdateDefaultUpdatedAt = userprofileDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userprofileDescID is the schema descriptor for id field.
	userprofileDescID := userprofileFields[0].Descriptor()
	// userprofile.DefaultID holds the default value on creation for the id field.
	userprofile.DefaultID = userprofileDescID.Default.(func() uuid.UUID)
	velogconfigFields := schema.VelogConfig{}.Fields()
	_ = velogconfigFields
	// velogconfigDescCreatedAt is the schema descriptor for created_at field.
	velogconfigDescCreatedAt := velogconfigFields[3].Descriptor()
	// velogconfig.DefaultCreatedAt holds the default value on creation for the created_at field.
	velogconfig.DefaultCreatedAt = velogconfigDescCreatedAt.Default.(func() time.Time)
	// velogconfigDescUpdatedAt is the schema descriptor for updated_at field.
	velogconfigDescUpdatedAt := velogconfigFields[4].Descriptor()
	// velogconfig.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	velogconfig.DefaultUpdatedAt = velogconfigDescUpdatedAt.Default.(func() time.Time)
	// velogconfig.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	velogconfig.UpdateDefaultUpdatedAt = velogconfigDescUpdatedAt.UpdateDefault.(func() time.Time)
	// velogconfigDescID is the schema descriptor for id field.
	velogconfigDescID := velogconfigFields[0].Descriptor()
	// velogconfig.DefaultID holds the default value on creation for the id field.
	velogconfig.DefaultID = velogconfigDescID.Default.(func() uuid.UUID)
}
