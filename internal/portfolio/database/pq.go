package database

import (
	"context"
	"portfolio/internal/portfolio"

	"gorm.io/gorm"
)

type PortfolioDb struct {
}

func Init() portfolio.DB {
	return &PortfolioDb{}
}

func (pdb *PortfolioDb) CreatePortfolio(ctx context.Context, p *portfolio.Portfolio) (*portfolio.Portfolio, error) {
	db := ctx.Value("surround").(map[string]interface{})["pgClient"].(*gorm.DB)
	tx := db.Create(p)
	return p, tx.Error
}
func (pdb *PortfolioDb) GetPortfolio(ctx context.Context, id float64) (*portfolio.Portfolio, error) {
	db := ctx.Value("surround").(map[string]interface{})["pgClient"].(*gorm.DB)
	p := &portfolio.Portfolio{}
	tx := db.Where("user_id=?", id).Preload("Entries").Find(p)
	return p, tx.Error
}
func (pdb *PortfolioDb) AddEntry(ctx context.Context, e *portfolio.Entry) (*portfolio.Entry, error) {
	db := ctx.Value("surround").(map[string]interface{})["pgClient"].(*gorm.DB)
	tx := db.Create(e)
	return e, tx.Error
}
