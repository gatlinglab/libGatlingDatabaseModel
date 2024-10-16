package dbModel

import "database/sql"

type IWJDatabase interface {
	Connect() error
	Close()
	ExecSql(sql string) (sql.Result, error)
	Query(sql string) (*sql.Rows, error)
}
