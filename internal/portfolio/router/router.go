package router

import (
	"portfolio/internal/portfolio"
	"portfolio/util"
	"portfolio/util/response"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Http struct {
	serv portfolio.Service
}

func Route(ser portfolio.Service, g *echo.Group, m ...echo.MiddlewareFunc) {
	h := &Http{
		serv: ser,
	}

	portfolioGrp := g.Group("/portfolio", m...)
	portfolioGrp.GET("/", h.GetPortfolio)
	portfolioGrp.POST("/:id/", h.AddEntry)
	portfolioGrp.POST("/", h.AddPortfolio)
}

func (h *Http) GetPortfolio(c echo.Context) error {
	user, err := h.serv.GetPortfolio(util.ToContextService(c))

	if err != nil {
		return response.RespondError(c, err)
	}
	return response.Respond(c, user)
}

func (h *Http) AddPortfolio(c echo.Context) error {
	user, err := h.serv.AddPortfolio(util.ToContextService(c))

	if err != nil {
		return response.RespondError(c, err)
	}
	return response.Respond(c, user)
}

func (h *Http) AddEntry(c echo.Context) error {
	e := &portfolio.Entry{}
	if err := c.Bind(e); err != nil {
		return response.RespondError(c, err)
	}
	e.PortfolioID, _ = strconv.Atoi(c.ParamValues()[0])
	user, err := h.serv.AddEntry(util.ToContextService(c), e)
	if err != nil {
		return response.RespondError(c, err)
	}
	return response.Respond(c, user)
}
