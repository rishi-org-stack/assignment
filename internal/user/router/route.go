package router

import (
	user "portfolio/internal/user"
	"portfolio/util"
	"portfolio/util/response"

	"github.com/labstack/echo/v4"
)

type Http struct {
	uSer user.Service
}

func Route(g *echo.Group, userService user.Service, m ...echo.MiddlewareFunc) {
	h := &Http{
		uSer: userService,
	}
	grpUser := g.Group("/user", m...)
	grpUser.GET("/", h.getById)

}

func (h *Http) getById(c echo.Context) error {

	user, err := h.uSer.GetUserByID(util.ToContextService(c))

	if err != nil {
		return response.RespondError(c, err)
	}
	return response.Respond(c, user)
}
