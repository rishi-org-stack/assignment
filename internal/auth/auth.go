package auth

import (
	"context"
	"net/http"
	"portfolio/internal/user"
	utilOTP "portfolio/util/auth"
	"portfolio/util/config"
	apiError "portfolio/util/error"
	apiRes "portfolio/util/response"
)

const (
	source                  = "AUTH"
	AUTH_INSERT_ERROR       = source + "_INSERT_ERROR"
	AUTH_SERVER_ERROR       = source + "_SERVER_ERROR"
	AUTH_BAD_REQUEST        = source + "_BAD_REQUEST"
	AUTH_OTP_INSERT_ERROR   = source + "_OTP_INSERT_ERROR"
	AUTH_UNAUTHORIZED_ERROR = source + "_INSERT_ERROR"
)

type AuthService struct {
	AuthData    DB
	JwtSer      TokenGenratorInterface
	Config      *config.Env
	UserService user.Service
}

var OTP string

func Init(db DB, js TokenGenratorInterface, us user.Service, config *config.Env) Service {
	return &AuthService{
		AuthData:    db,
		JwtSer:      js,
		Config:      config,
		UserService: us,
	}
}

func (authSer AuthService) HandleAuth(ctx context.Context, atr *AuthRequest) (*apiRes.Response, apiError.ApiErrorInterface) {
	// atr := &AuthRequest{
	// 	Email: "rishi@gmail.com",
	// }
	if authSer.Config.Mode == "dev" {
		res, err := authSer.AuthData.FindOrInsert(ctx, atr)
		if err != nil {
			return &apiRes.Response{}, apiError.ApiError{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
				Code:    AUTH_INSERT_ERROR,
			}
		}
		return &apiRes.Response{
			Status:  http.StatusOK,
			Message: "Email authenticated",
			Data: &AuthResponse{
				ID:  res.ID,
				OTP: "666666",
			},
		}, nil
	}
	res, err := authSer.AuthData.FindOrInsert(ctx, atr)
	if err != nil {
		return &apiRes.Response{}, apiError.ApiError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Code:    AUTH_INSERT_ERROR,
		}
	}
	otp := utilOTP.GenrateOtp(authSer.Config.OTPExpiry)
	if err := otp.Set(res.ID); err != nil {
		return &apiRes.Response{}, apiError.ApiError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Code:    AUTH_OTP_INSERT_ERROR,
		}
	}
	// OTP = otp.Otp
	return &apiRes.Response{
		Status:  http.StatusOK,
		Message: "Email authenticated",
		Data: &AuthResponse{
			ID:  res.ID,
			OTP: otp.Otp,
		},
	}, nil

}
func (authSer AuthService) Verify(ctx context.Context, otpReq *VerifyRequest) (*apiRes.Response, apiError.ApiErrorInterface) {
	otpGiven := "666666"
	if authSer.Config.Mode != "dev" {

		otp := &utilOTP.OTP{}
		Otp, err := otp.Get(otpReq.ID)
		if err != nil {
			return &apiRes.Response{}, apiError.ApiError{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
				Code:    AUTH_BAD_REQUEST,
			}
		}
		otpGiven = Otp
	}
	if otpGiven != otpReq.OTP {
		return &apiRes.Response{}, apiError.ApiError{
			Status:  http.StatusUnauthorized,
			Message: "otp doesn't matches",
			Code:    AUTH_UNAUTHORIZED_ERROR,
		}
	}
	id := otpReq.ID
	req, err := authSer.AuthData.GetRequest(ctx, otpReq.ID)
	if err != nil {
		return &apiRes.Response{}, apiError.ApiError{
			Status:  http.StatusBadRequest,
			Message: "pls try after some time",
			Code:    AUTH_SERVER_ERROR,
		}
	}
	user, err := authSer.UserService.FindOrCreateUser(ctx, req.Email)
	if err != nil {
		return &apiRes.Response{}, apiError.ApiError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Code:    AUTH_SERVER_ERROR,
		}
	}
	id = user.ID

	token, err := authSer.createToken(id, req.Email)
	if err != nil {
		return &apiRes.Response{}, apiError.ApiError{
			Status:  http.StatusInternalServerError,
			Message: "pls try after some time",
			Code:    AUTH_SERVER_ERROR,
		}
	}
	return &apiRes.Response{
		Status:  http.StatusOK,
		Message: "otp  verified",
		Data:    token,
	}, nil
}

func (s AuthService) createToken(id int, email string) (string, error) {
	token, err := s.JwtSer.GenrateToken(id, email)
	if err != nil {
		return "", err
	}
	return token, nil
}
func (ar AuthService) GetRequestByID(ctx context.Context, id int) (*AuthRequest, error) {
	authR, err := ar.AuthData.GetRequest(ctx, 2)
	if err != nil {
		return &AuthRequest{}, nil
	}
	return authR, nil
}
