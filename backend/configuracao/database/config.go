package database

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	config "github.com/isaqueveras/ufc-projeto-banco-dados-postgresql/backend/configuracao"
	_ "github.com/lib/pq"
)

// DBTransacao usado para agregar as conexões
type DBTransacao struct {
	postgres *sql.Tx
	Builder  squirrel.StatementBuilderType
}

// AbrirTransacao inicializa uma conexão com o banco de dados
func AbrirTransacao() (*DBTransacao, error) {
	var (
		t           = &DBTransacao{}
		db          *sql.DB
		transaction *sql.Tx
		err         error
	)

	if db, err = sql.Open(config.Obter().Database.Driver, config.Obter().Database.Url); err != nil {
		return t, err
	}

	defer db.Close()

	if transaction, err = db.BeginTx(context.Background(), &sql.TxOptions{
		Isolation: 0,
		ReadOnly:  false,
	}); err != nil {
		return nil, err
	}

	t.postgres = transaction
	t.Builder = squirrel.StatementBuilder.
		PlaceholderFormat(squirrel.Dollar).
		RunWith(t.postgres)

	return t, nil
}

// Commit commit pending transactions for all open databases
func (t *DBTransacao) Commit() (erro error) {
	return t.postgres.Commit()
}

// Rollback rollback transaction pending for all open databases
func (t *DBTransacao) Rollback() {
	_ = t.postgres.Rollback()
}
