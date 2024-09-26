package handlers

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"primijenjena-informatika-dev/ent"
	"primijenjena-informatika-dev/internal/services"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func NewCodespaceHandler(db *ent.Client, c echo.Context) error {
	code := c.QueryParam("code")
	urlState := c.QueryParam("state")
	stateCookie, err := c.Cookie("oauth_state")
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid OAuth2 state",
		})
	}

	if urlState != stateCookie.Value {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "OAuth2 state mismatch",
		})
	}

	webUrl, err := services.OpenCodespace(db, code, c.RealIP())
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to create codespace",
		})
	}

	c.SetCookie(&http.Cookie{
		Name:     "oauth_state",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Now().Add(-time.Hour),
	})

	return c.Redirect(http.StatusTemporaryRedirect, webUrl)
}

func OauthHandler(c echo.Context) error {
	oauthState, err := randomBytes(32)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to generate oauth state",
		})
	}

	c.SetCookie(&http.Cookie{
		Name:     "oauth_state",
		Value:    base64.URLEncoding.EncodeToString(oauthState),
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Now().Add(time.Minute * 5),
	})

	authUrl := fmt.Sprintf("https://github.com/login/oauth/authorize?client_id=%s&state=%s&redirect_uri=%s&scope=repo admin:org user codespace", viper.GetString("github_client_id"), base64.URLEncoding.EncodeToString(oauthState), viper.GetString("redirect_uri"))

	return c.Redirect(http.StatusTemporaryRedirect, authUrl)
}

func randomBytes(n int) ([]byte, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return nil, err
	}
	return bytes, nil
}
