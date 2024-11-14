package dbModel

import (
	"database/sql"
	"time"
)

type IWJDatabase interface {
	Connect(sslConfig string) error
	Close()
	GetDBHandler() *sql.DB
	GetDatabaseVersion() (string, error)
	GetDatabaseType() DBMWJDatabaseType
	//GetLastError() error
	ExecSql(sql string, args ...any) (sql.Result, error)
	Query(sql string) (*sql.Rows, error)
	CheckTableExists(tableName string) bool
	//CreateTableIfNotExists(tableType DBMWJTableType, tableName string) error
	DropTableIfExists(tableName string) error
	SetTimeOutSeconds(time.Duration)
}
