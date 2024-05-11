package database

import (
	"context"
	"fmt"
	"github.com/exaring/otelpgx"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel"
	"time"
)

const (
	defaultQueryExecMode = "cache_describe"
	sslMode              = "disable"
	defaultTimeout       = 5 * time.Second
)

type DBCfg struct {
	Host            string        `yaml:"host" env:"HOST" validate:"required"`
	Port            uint16        `yaml:"port" env:"PORT" validate:"required"`
	DB              string        `yaml:"db" env:"DB" validate:"required"`
	User            string        `yaml:"user" env:"USER" validate:"required"`
	Password        string        `yaml:"password" env:"PASSWORD" validate:"required"`
	MaxIdleLifetime time.Duration `yaml:"max_idle_lifetime" env:"MAX_IDLE_LIFETIME" env-default:"10s" validate:"required"`
	MaxConnLifetime time.Duration `yaml:"max_conn_lifetime" env:"MAX_CONN_LIFETIME" env-default:"5m" validate:"required"`
	MaxConnAmount   int32         `yaml:"max_conn_amount" env:"MAX_CONN_AMOUNT" env-default:"30" validate:"required"`
	MinConnAmount   int32         `yaml:"min_conn_amount" env:"MIN_CONN_AMOUNT" env-default:"5" validate:"required"`
}

func (cfg DBCfg) String() string {
	return fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=%s default_query_exec_mode=%s",
		cfg.Host,
		cfg.Port,
		cfg.DB,
		cfg.User,
		cfg.Password,
		sslMode,
		defaultQueryExecMode,
	)
}

func NewPool(cfg DBCfg) (*pgxpool.Pool, error) {
	dbConfig, err := pgxpool.ParseConfig(cfg.String())
	if err != nil {
		return nil, err
	}
	dbConfig.MaxConnIdleTime = cfg.MaxIdleLifetime
	dbConfig.MaxConnLifetime = cfg.MaxConnLifetime
	dbConfig.MaxConns = cfg.MaxConnAmount
	dbConfig.MinConns = cfg.MinConnAmount
	dbConfig.ConnConfig.Tracer = otelpgx.NewTracer(otelpgx.WithTracerProvider(otel.GetTracerProvider()))

	dbpool, err := pgxpool.NewWithConfig(context.Background(), dbConfig)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	if err := dbpool.Ping(ctx); err != nil {
		return nil, errors.Wrap(err, "failed to ping sql")
	}

	return dbpool, nil
}
