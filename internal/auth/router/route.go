package router

import (
	"context"
	"fmt"
	"net/http"
	"portfolio/internal/auth"
	"portfolio/util"
	apiRes "portfolio/util/response"

	"github.com/labstack/echo/v4"
)

type Http struct {
	serv auth.Service
}

func Route(ser auth.Service, g *echo.Group, m ...echo.MiddlewareFunc) {
	h := &Http{
		serv: ser,
	}
	grpAuth := g.Group("/auth", m...)
	grpAuth.POST("/", h.ok)
	grpAuth.POST("/verify", h.verify)
}
func (h *Http) ok(c echo.Context) error {
	ctx := context.WithValue(context.Background(), "pgClient", c.Get("pgClient"))
	atr := &auth.AuthRequest{}
	if err := c.Bind(atr); err != nil {
		fmt.Println("log of ok router")
		return err
	}
	fmt.Println("no error")
	res, err := h.serv.HandleAuth(ctx, atr)
	if err != nil {
		return apiRes.RespondError(c, err)
	}
	return c.JSON(http.StatusOK, res)
}

func (h *Http) verify(c echo.Context) error {
	var verifyReq = &auth.VerifyRequest{}
	if err := c.Bind(verifyReq); err != nil {
		return err
	}
	res, err := h.serv.Verify(util.ToContextService(c), verifyReq)
	if err != nil {
		return apiRes.RespondError(c, err)
	}
	return apiRes.Respond(c, res)
}
