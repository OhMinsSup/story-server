package controllers

import (
	"fmt"
	"github.com/OhMinsSup/story-server/database/models"
	"github.com/OhMinsSup/story-server/helpers"
	"github.com/OhMinsSup/story-server/helpers/social"
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/github"
	"github.com/jinzhu/gorm"
	"net/http"
	"strings"
)

func SocialRedirect(ctx *gin.Context) {
	provider := ctx.Param("provider")
	next := ctx.Query("next")

	providerType := []string{
		"facebook",
		"github",
		"google",
	}

	if !strings.Contains(strings.Join(providerType, ","), provider) {
		ctx.AbortWithError(http.StatusBadRequest, helpers.ErrorProviderValided)
		return
	}

	loginUrl := social.GenerateSocialLink(provider, next)
	ctx.Redirect(http.StatusMovedPermanently, loginUrl)
}

func GithubCallback(ctx *gin.Context) {
	code := ctx.Query("code")
	if code == "" {
		ctx.AbortWithError(http.StatusBadRequest, helpers.ErrorNotFound)
		return
	}

	provider := "github"
	accessToken := social.GetGithubAccessToken(code)
	profile := social.GetGithubProfile(accessToken)

	db := ctx.MustGet("db").(*gorm.DB)

	var data models.SocialAccount
	if err := db.Where(&models.SocialAccount{
		SocialId: fmt.Sprint(profile.ID),
		Provider: provider,
	}).First(&data).Error; err != nil {
		panic(err)
		return
	}

	ctx.Set("profile", profile)
	ctx.Set("social", data)
	ctx.Set("accessToken", accessToken)
	ctx.Set("provider", provider)
	ctx.Next()
	return
}



func SocialCallback(ctx *gin.Context) {
	profile := ctx.MustGet("profile").(*github.User)
	social := ctx.MustGet("social").(*models.SocialAccount)
	accessToken := ctx.MustGet("accessToken").(string)
	provider := ctx.MustGet("provider").(string)

	if profile == nil || accessToken == "" {
		ctx.AbortWithError(http.StatusForbidden, helpers.ErrorForbidden)
		return
	}

	db := ctx.MustGet("db").(*gorm.DB)
	var user models.User
	if social != nil {
		if err := db.Where("user_id = ?", social.ID).First(&user).Error; err != nil {
			ctx.AbortWithError(http.StatusNotFound, helpers.ErrorUserIsMissing)
			return
		}

		tokens := user.GenerateUserToken(db)
		ctx.SetCookie("access_token", tokens["accessToken"].(string), 60*60*24, "/", "", false, true)
		ctx.SetCookie("refresh_token", tokens["refreshToken"].(string), 60*60*24*30, "/", "", false, true)

		redirectUrl := "http://localhost:3000"
		next := ""
		ctx.Redirect(http.StatusMovedPermanently, redirectUrl+next)
		return
	}

	if profile.Email != nil {
		db.Where("email = ?", profile.Email).First(&user)
	}

	if &user != nil {
		tokens := user.GenerateUserToken(db)
		ctx.SetCookie("access_token", tokens["accessToken"].(string), 60*60*24, "/", "", false, true)
		ctx.SetCookie("refresh_token", tokens["refreshToken"].(string), 60*60*24*30, "/", "", false, true)
		redirectUrl := "https://localhost:3000/"
		ctx.Redirect(http.StatusMovedPermanently, redirectUrl)
		return
	}

	payload := helpers.JSON{
		"profile":     profile,
		"provider":    provider,
		"accessToken": accessToken,
	}

	registerToken, err := helpers.GenerateRegisterToken(payload, "")
	if err != nil {
		ctx.AbortWithError(http.StatusConflict, err)
		return
	}

	ctx.SetCookie("register_token", registerToken, 60*60, "/", "", false, true)
	redirectUrl := "http://localhost:3000/register?social=1"
	ctx.Redirect(http.StatusMovedPermanently, redirectUrl)
}

func GoogleCallback(ctx *gin.Context) {
	ctx.Next()
}

func FacebookCallback(ctx *gin.Context) {
	ctx.Next()
}
