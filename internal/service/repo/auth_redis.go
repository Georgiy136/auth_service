package repo

import (
	"context"
	"fmt"
	"github.com/Georgiy136/auth_service/internal/models"
	jsoniter "github.com/json-iterator/go"
	"github.com/redis/go-redis/v9"
)

type AuthRedis struct {
	Rdb *redis.Client
}

func NewAuthRedis(rdb *redis.Client) AuthDBStore {
	return &AuthRedis{
		Rdb: rdb,
	}
}

var key = "%s_%s"

func (db *AuthRedis) SaveUserSession(ctx context.Context, data models.LoginInfo) error {
	dataString, err := jsoniter.MarshalToString(data)
	if err != nil {
		return fmt.Errorf("SaveUserSession - jsoniter.MarshalToString: %w", err)
	}

	if err = db.Rdb.Set(ctx, fmt.Sprintf(key, data.UserID, data.SessionID), dataString, redis.KeepTTL).Err(); err != nil {
		return fmt.Errorf("SaveUserSession err: %v", err)
	}
	return nil
}

func (db *AuthRedis) GetUserSession(ctx context.Context, userID int, sessionID string) (*models.LoginInfo, error) {
	data, err := db.Rdb.Get(ctx, fmt.Sprintf(key, userID, sessionID)).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, fmt.Errorf("GetUserSession err: %w", err)
	}

	var result *models.LoginInfo
	if err = jsoniter.UnmarshalFromString(data, result); err != nil {
		return nil, fmt.Errorf("GetUserSession - jsoniter.UnmarshalFromString: %w", err)
	}

	return result, nil
}

func (db *AuthRedis) DeleteUserSession(ctx context.Context, userID int, sessionID string) error {
	if err := db.Rdb.Del(ctx, fmt.Sprintf(key, userID, sessionID)).Err(); err != nil {
		return fmt.Errorf("DeleteUserSession err: %v", err)
	}
	return nil
}
