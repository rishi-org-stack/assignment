package user

import (
	utilError "portfolio/util/error"
	"portfolio/util/response"

	// "askUs/v1/util/response"
	"context"
	"net/http"
)

const (
	source = "USER"

	ID_DECODE_ERROR           = source + "_ERROR_GEETING_ID"
	USER_GET_ERROR            = source + "_GET_ERROR"
	USER_DOCTOR_CREATE_ERROR  = source + "_DOCTOR_CREATE_ERROR"
	USER_PATIENT_CREATE_ERROR = source + "_PATIENT_CREATE_ERROR"
	USER_COPY_ERROR           = source + "_COPY_ERROR"
)

type (
	UserService struct {
		UserData DB
		// AuthService auth.Service
	}
)

func Init(db DB) Service {
	return &UserService{
		UserData: db,
		// AuthService: authser,
	}

}

func (s UserService) FindOrCreateUser(ctx context.Context, email string) (*Usr, utilError.ApiErrorInterface) {

	doc, err := s.UserData.FindOrCreateUser(ctx, email)

	if err != nil {
		return &Usr{}, utilError.ApiError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Code:    USER_DOCTOR_CREATE_ERROR,
		}
	}
	return doc, nil
}

func (s UserService) GetUserByID(ctx context.Context) (*response.Response, utilError.ApiErrorInterface) {
	id := ctx.Value("surround").(map[string]interface{})["id"].(float64)
	doc, err := s.UserData.GetUserByID(ctx, id)
	if err != nil {
		return &response.Response{}, utilError.ApiError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Code:    USER_GET_ERROR,
		}
	}
	return &response.Response{
		Status:  http.StatusOK,
		Message: "User retrieve successfull",
		Data:    doc,
	}, nil
}
