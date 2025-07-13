package token_generate

import (
	"github.com/Georgiy136/auth_service/internal/service/token_generate/tokens"
)

type IssueTokensService struct {
	RefreshToken tokens.RefreshTokenStore
	AccessToken  tokens.AccessTokenStore
}

func NewIssueTokensService(
	refreshToken tokens.RefreshTokenStore,
	accessToken tokens.AccessTokenStore,
) *IssueTokensService {
	return &IssueTokensService{
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
	}
}
