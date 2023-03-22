package postgres

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/perfectogo/template-service-with-kafka/internal/entity"
	configpkg "github.com/perfectogo/template-service-with-kafka/internal/pkg/config"
)

// PostgresDB ...
type PostgresDB struct {
	*pgxpool.Pool
	Sq *Squirrel
}

// New provides PostgresDB struct init
func New(config *configpkg.Config) (*PostgresDB, error) {

	db := PostgresDB{Sq: NewSquirrel()}

	if err := db.connect(config); err != nil {
		return nil, err
	}

	return &db, nil
}

func (p *PostgresDB) configToStr(config *configpkg.Config) string {
	var conn []string

	if len(config.DB.Host) != 0 {
		conn = append(conn, "host="+config.DB.Host)
	}

	if len(config.DB.Port) != 0 {
		conn = append(conn, "port="+config.DB.Port)
	}

	if len(config.DB.User) != 0 {
		conn = append(conn, "user="+config.DB.User)
	}

	if len(config.DB.Password) != 0 {
		conn = append(conn, "password="+config.DB.Password)
	}

	if len(config.DB.Name) != 0 {
		conn = append(conn, "dbname="+config.DB.Name)
	}

	if len(config.DB.Sslmode) != 0 {
		conn = append(conn, "sslmode="+config.DB.Sslmode)
	}

	return strings.Join(conn, " ")
}

func (p *PostgresDB) connect(config *configpkg.Config) error {
	pgxpoolConfig, err := pgxpool.ParseConfig(p.configToStr(config))
	if err != nil {
		return fmt.Errorf("unable to parse database config: %w", err)
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), pgxpoolConfig)
	if err != nil {
		return fmt.Errorf("unable to connect database config: %w", err)
	}

	p.Pool = pool

	return nil
}

func (p *PostgresDB) Close() {
	p.Pool.Close()
}

func (p *PostgresDB) Error(err error) error {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23505":
			return entity.ErrorConflict
		}
	}
	if err == pgx.ErrNoRows {
		return entity.ErrorNotFound
	}
	return err
}

func (p *PostgresDB) ErrSQLBuild(err error, message string) error {
	return fmt.Errorf("error during sql build, %s: %w", message, err)
}
