package interfaces
import (
	"context"
	"database/sql"
)

type DBStore interface{ 
	Begin()  (*sql.Tx, error)
	BeginTx(context.Context, *sql.TxOptions)  (*sql.Tx, error)
	Prepare(query string) (*sql.Stmt, error)
    QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
    ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
    Close() error
}