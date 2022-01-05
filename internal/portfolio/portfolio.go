package portfolio

import (
	"context"
	"fmt"
	"net/http"
	apiError "portfolio/util/error"
	apiRes "portfolio/util/response"
)

type PortfolioService struct {
	repo DB
}

func Init(repo DB) Service {
	return &PortfolioService{
		repo: repo,
	}
}

func (ps *PortfolioService) GetPortfolio(ctx context.Context) (*apiRes.Response, apiError.ApiErrorInterface) {
	id := ctx.Value("surround").(map[string]interface{})["id"].(float64)
	p, err := ps.repo.GetPortfolio(ctx, id)
	if err != nil {
		return &apiRes.Response{}, apiError.ApiError{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return &apiRes.Response{
		Status:  http.StatusOK,
		Message: "get portfolio done",
		Data:    p,
	}, nil
}
func (ps *PortfolioService) AddPortfolio(ctx context.Context) (*apiRes.Response, apiError.ApiErrorInterface) {
	uid := ctx.Value("surround").(map[string]interface{})["id"].(float64)
	p := &Portfolio{}
	p.UserID = fmt.Sprint(int(uid))
	p, err := ps.repo.CreatePortfolio(ctx, p)
	if err != nil {
		return &apiRes.Response{}, apiError.ApiError{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return &apiRes.Response{
		Status:  http.StatusOK,
		Message: "add portfolio done",
		Data:    p,
	}, nil
}
func (ps *PortfolioService) AddEntry(ctx context.Context, e *Entry) (*apiRes.Response, apiError.ApiErrorInterface) {
	p, err := ps.repo.AddEntry(ctx, e)
	if err != nil {
		return &apiRes.Response{}, apiError.ApiError{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return &apiRes.Response{
		Status:  http.StatusOK,
		Message: "add entry done",
		Data:    p,
	}, nil
}
