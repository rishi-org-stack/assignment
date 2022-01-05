package user

//TODO: request service is almost done need:
//TODO: a route to get all followed doctors following doctors followed by patients
//TODO: needa rbac for user request module Done
//TODO: need a  new caching system
import (
	"context"
	utilError "portfolio/util/error"
	"portfolio/util/response"
)

type (
	DB interface {
		FindOrCreateUser(ctx context.Context, email string) (*Usr, error)
		GetUserByID(ctx context.Context, id float64) (*Usr, error)
	}
	Service interface {
		FindOrCreateUser(ctx context.Context, email string) (*Usr, utilError.ApiErrorInterface)
		GetUserByID(ctx context.Context) (*response.Response, utilError.ApiErrorInterface)
	}

	//TODO:User ID needs to be of type Object Id
	Usr struct {
		ID int `json:"id" gorm:"primary"`
		Info
	}
	//FD

	Info struct {
		Email string `json:"email"`
		Phone string `json:"phone"`
		Name  string `json:"name" gorm:"not null"`

		// Address
	}
)

const Schema = "usr"

func (Usr) TableName() string {
	return Schema + ".usrs"
}
