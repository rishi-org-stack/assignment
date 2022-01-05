package portfolio

import (
	"context"
	apiError "portfolio/util/error"
	apiRes "portfolio/util/response"
)

type (
	Service interface {
		GetPortfolio(context.Context) (*apiRes.Response, apiError.ApiErrorInterface)
		AddPortfolio(context.Context) (*apiRes.Response, apiError.ApiErrorInterface)
		AddEntry(context.Context, *Entry) (*apiRes.Response, apiError.ApiErrorInterface)
	}

	DB interface {
		CreatePortfolio(context.Context, *Portfolio) (*Portfolio, error)
		GetPortfolio(context.Context, float64) (*Portfolio, error)
		AddEntry(context.Context, *Entry) (*Entry, error)
	}

	Portfolio struct {
		ID      int     `json:"id"  gorm:"primaryKey"`
		UserID  string  `json:"userID" gorm:"not null"`
		Entries []Entry `json:"entries"`
	}
	Entry struct {
		ID             int        `json:"id"  gorm:"primaryKey"`
		CoinName       string     `json:"coinName"  gorm:"not null"`
		Amount         int        `json:"amount"  gorm:"not null"`
		Price          float64    `json:"price"  gorm:"not null"`
		TransactionFee float64    `json:"transactionFee"  gorm:"not null"`
		PortfolioID    int        `json:"portfolioID"`
		Portfolio      *Portfolio `json:"portfolio"`
	}
)

const (
	schema = "portfolio"
)

func (Portfolio) TableName() string {
	return schema + ".portfolios"
}

func (Entry) TableName() string {
	return schema + ".entries"
}
