package tokens

import "github.com/Georgiy136/auth_service/internal/models"

type (
	AccessTokenStore interface {
		New(accessTokenPayload models.AccessTokenPayload) (string, error)
		Parse(accessToken string) (*models.AccessTokenPayload, error)
	}

	RefreshTokenStore interface {
		New() (string, error)
		Parse(refreshToken string) error
	}
)
