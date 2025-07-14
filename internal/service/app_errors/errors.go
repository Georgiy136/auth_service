package app_errors

import "errors"

var (
	DecodeAndDecryptTokenError = errors.New("decode and decrypt token error")
	ParseJwtTokenError         = errors.New("parse jwt token error")
	TokenIsExpiredError        = errors.New("token is expired")
	SessionUserNotFoundError   = errors.New("user session not found")
	UserAgentNotMatchError     = errors.New("User-Agent does not match in db")
	RefreshTokenNotValidError  = errors.New("refresh token not valid")
)
