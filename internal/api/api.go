package api

import (
	"portfolio/internal/auth"
	amdb "portfolio/internal/auth/databases/psql"
	authR "portfolio/internal/auth/router"
	"portfolio/internal/portfolio"

	// umdb "askUs/v1/package/user/databases/psql"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
	"os"
	pdb "portfolio/internal/portfolio/database"
	portfolioR "portfolio/internal/portfolio/router"
	"portfolio/internal/user"
	umdb "portfolio/internal/user/databases/psql"
	userR "portfolio/internal/user/router"
	jAuth "portfolio/util/auth"
	"portfolio/util/config"
	mid "portfolio/util/middleware"
)

type api struct {
	Client      *gorm.DB
	Version     string
	MiddleWares []echo.MiddlewareFunc
	Jwt         *jAuth.Auth
	Config      *config.Env
}

func Init(c *gorm.DB, jwt *jAuth.Auth, env *config.Env, m ...echo.MiddlewareFunc) *api {
	return &api{
		Client:      c,
		Version:     os.Getenv("VERSION"),
		MiddleWares: m,
		Jwt:         jwt,
		Config:      env,
	}
}
func (ap *api) Route(e *echo.Echo) {
	e.Use(mid.ConnectionMDB(ap.Client), mid.Logger())

	v1 := e.Group("/api/" + ap.Version)
	v1.GET("/ping", func(c echo.Context) error {
		return c.JSON(200, "pong")
	})
	userService := user.Init(&umdb.UserDb{})
	portfolioService := portfolio.Init(&pdb.PortfolioDb{})

	authService := auth.Init(amdb.AuthDb{}, ap.Jwt, userService, ap.Config)

	authR.Route(authService, v1, mid.ConnectionMDB(ap.Client))
	userR.Route(v1, userService, ap.MiddleWares...)
	portfolioR.Route(portfolioService, v1, ap.MiddleWares...)

}
