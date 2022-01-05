package psql

import (
	"context"
	"portfolio/internal/user"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserDb struct{}

func (udb *UserDb) FindOrCreateUser(ctx context.Context, email string) (*user.Usr, error) {
	// db := ctx.Value("pgClient").(*gorm.D
	db := ctx.Value("surround").(map[string]interface{})["pgClient"].(*gorm.DB)
	doc := &user.Usr{Info: user.Info{Email: email}}
	tx := db.Where("email=?", email).
		FirstOrCreate(doc)
	return doc, tx.Error
}

func (udb *UserDb) GetUserByID(ctx context.Context, id float64) (*user.Usr, error) {
	db := ctx.Value("surround").(map[string]interface{})["pgClient"].(*gorm.DB)
	doc := &user.Usr{}
	model := db.Model(doc)
	tx := model.Preload(clause.Associations).First(doc, "id=?", id)
	return doc, tx.Error
}
