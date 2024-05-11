package database

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

//go:generate minimock -g -s .go -i github.com/jackc/pgx/v5.Row -o ../../mocks/pgx
//go:generate minimock -g -s .go -i github.com/jackc/pgx/v5.Rows -o ../../mocks/pgx
//go:generate minimock -g -s .go -i github.com/jackc/pgx/v5.Tx -o ../../mocks/pgx

//go:generate minimock -g -s .go -i IDatabase -o ../../mocks/pkg/database

type IDatabase interface {
	IQuery
	IQueryRow
	IExec
	Begin(ctx context.Context) (pgx.Tx, error)
}

//go:generate minimock -g -s .go -i IExec -o ../../mocks/pkg/database
type IExec interface {
	Exec(context.Context, string, ...interface{}) (commandTag pgconn.CommandTag, err error)
}

type IQueryRow interface {
	QueryRow(context.Context, string, ...interface{}) pgx.Row
}

type IQuery interface {
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
}
