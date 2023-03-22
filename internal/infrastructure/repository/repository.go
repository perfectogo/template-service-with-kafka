package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"

	otlp_pkg "github.com/perfectogo/template-service-with-kafka/internal/pkg/otlp"
	"github.com/perfectogo/template-service-with-kafka/internal/pkg/postgres"
)

type BaseRepository struct {
	db        *postgres.PostgresDB
	repoName  string
	tableName string
}

func (br *BaseRepository) SetDB(db *postgres.PostgresDB) {
	br.db = db
}

func (br *BaseRepository) SetRepoName(repoName string) {
	br.repoName = repoName
}

func (br *BaseRepository) SetTableName(tableName string) {
	br.tableName = tableName
}

func (br *BaseRepository) DB() *postgres.PostgresDB {
	return br.db
}

func (br *BaseRepository) RepoName() string {
	return br.repoName
}

func (br *BaseRepository) TableName() string {
	return br.tableName
}

func (br *BaseRepository) TableNameAlias(alias string) string {
	return fmt.Sprintf("%s as %s", br.tableName, alias)
}

func (br *BaseRepository) SpanName(methodName string) string {
	return fmt.Sprintf("%s.%s", br.repoName, methodName)
}

func (br *BaseRepository) Tracing(ctx context.Context, methodName string) (context.Context, otlp_pkg.Span) {
	return otlp_pkg.Start(ctx, br.RepoName(), br.SpanName(methodName))
}

func (br *BaseRepository) TxRollback(ctx context.Context, tx postgres.Tx, err error) error {
	if txErr := tx.Rollback(ctx); txErr != nil {
		if err.Error() != pgx.ErrTxClosed.Error() {
			return fmt.Errorf("%w: %v", err, txErr)
		}
	}
	return err
}
