package dbModel

import "database/sql"

type IWJDatabase interface {
	Connect() error
	Close()
	GetDatabaseVersion() string
	GetDatabaseType() DBMWJDatabaseType
	//GetLastError() error
	ExecSql(sql string) (sql.Result, error)
	Query(sql string) (*sql.Rows, error)
	CheckTableExists(tableName string) bool
	//CreateTableIfNotExists(tableType DBMWJTableType, tableName string) error
	DropTableIfExists(tableName string) error
}
