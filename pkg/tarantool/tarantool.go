package tarantool

/*
import (
	"fmt"
	"github.com/Georgiy136/auth_service/config"
	"github.com/tarantool/go-tarantool"
)

type Tarantool struct {
	Conn *tarantool.Connection
}

func NewTarantool(cfg config.Tarantool) (*Tarantool, error) {
	conn, err := tarantool.Connect(fmt.Sprintf("%s:%d", cfg.Host, cfg.Port), tarantool.Opts{
		User: cfg.User,
		Pass: cfg.Password,
	})
	if err != nil {
		return nil, fmt.Errorf("connecting to tarantool error: %w", err)
	}

	return &Tarantool{
		Conn: conn,
	}, nil
}*/
