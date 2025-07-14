package repo

/*
import (
	"context"
	"fmt"
	"github.com/Georgiy136/auth_service/internal/models"
	pkg_conn "github.com/Georgiy136/auth_service/pkg/tarantool"
	jsoniter "github.com/json-iterator/go"
	"github.com/tarantool/go-tarantool"
)

type AuthTarantool struct {
	conn *tarantool.Connection
}

func NewAuthTarantool(db *pkg_conn.Tarantool) AuthDBStore {
	return &AuthTarantool{
		conn: db.Conn,
	}
}

func (db *AuthTarantool) SaveUserSession(ctx context.Context, data models.LoginInfo) error {
	const cp = "save_session"

	resp, err := db.conn.Call(cp, []interface{}{ctx, data.UserID, data.SessionID, data.Token, data.UserAgent, data.IpAddress})
	if err != nil {
		return fmt.Errorf("SaveUserSession err: %v, resp.Error", err)
	}
	if resp == nil {
		return fmt.Errorf("SaveUserSession resp is nil")
	}
	if resp.Code != 204 {
		return fmt.Errorf("SaveUserSession - resp.Code not 204 err: %v", resp)
	}

	return nil
}

func (db *AuthTarantool) GetUserSession(ctx context.Context, userID int, sessionID string) (*models.LoginInfo, error) {
	const cp = "get_session"

	resp, err := db.conn.Call(cp, []interface{}{ctx, userID, sessionID})
	if err != nil {
		return nil, fmt.Errorf("GetUserSession err: %v", err)
	}

	var result models.LoginInfo
	if err = jsoniter.Unmarshal(resp.Data[0].([]byte), &result); err != nil {
		return nil, fmt.Errorf("GetUserSession - unmarshal resp error: %v", err)
	}

	return &result, nil
}

func (db *AuthTarantool) DeleteUserSession(ctx context.Context, userID int, sessionID string) error {
	const cp = "del_session"

	_, err := db.conn.Call(cp, []interface{}{ctx, userID, sessionID})
	if err != nil {
		return fmt.Errorf("DeleteUserSession err: %v", err)
	}

	return nil
}
*/
