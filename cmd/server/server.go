package server

import (
	"embed"
	"fmt"
	"primijenjena-informatika-dev/internal/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"golang.org/x/crypto/acme/autocert"

	_ "github.com/mattn/go-sqlite3"
)

func CreateServer(publicFS embed.FS) *echo.Echo {
	db := CreateDatabaseClient()

	e := echo.New()

	e.AutoTLSManager.Cache = autocert.DirCache("/var/www/.cache")
	e.Use(middleware.Logger())
	// e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	if viper.GetBool("enable_tls") {
		e.Use(middleware.Secure())
	}

	fs := echo.MustSubFS(publicFS, "public")
	e.StaticFS("/", fs)
	e.GET("/new", func(c echo.Context) error {
		return handlers.NewCodespaceHandler(db, c)
	})
	e.GET("/oauth", func(c echo.Context) error {
		return handlers.OauthHandler(c)
	})

	return e
}

func StartServer(server *echo.Echo) {
	if viper.GetBool("enable_tls") {
		server.Logger.Fatal(server.StartAutoTLS(fmt.Sprintf(":%d", viper.GetInt("port"))))
	} else {
		server.Logger.Fatal(server.Start(fmt.Sprintf(":%d", viper.GetInt("port"))))
	}
}
