package manager

import (
	"database/sql"
	"employeeleave/config"
	"fmt"
)

type InfraManager interface {
	Conn() *sql.DB
}

type infraManager struct {
	db  *sql.DB
	cfg *config.Config
}

func (i *infraManager) initDb() error {
	dsn := fmt.Sprintf("host=%s port=%s database=%s user=%s password=%s sslmode=disable", i.cfg.Host, i.cfg.Port, i.cfg.Name, i.cfg.User, i.cfg.Password)
	db, err := sql.Open(i.cfg.Driver, dsn)
	if err != nil {
		return err
	}
	i.db = db
	return nil
}

func (i *infraManager) Conn() *sql.DB {
	return i.db
}

func NewInfraManager(cfg *config.Config) (InfraManager, error) {
	conn := &infraManager{
		cfg: cfg,
	}
	err := conn.initDb()
	if err != nil {
		return nil, err
	}
	return conn, nil
}
