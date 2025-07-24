package interfaces
import "database/sql"

type DBStore interface{ 
	Begin()  (*sql.Tx, error)
	Prepare(query string) (*sql.Stmt, error) 
}