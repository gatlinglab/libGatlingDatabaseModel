package dbModel

import (
	"database/sql"
	"time"
)

// id primary, Key string, ValueStr string, ValueInt int, ValueFloat float64,date default now;
type CWJDBTableIDKVD1 struct {
	ID         int64
	Key        string
	ValueStr   string
	ValueInt   int64
	ValueFloat float64
	Date1      time.Time
}

const CWJDBTD_IDKVD1_CreateTableSqlite = ` Create Table %s(
	ID                    BigInt         NOT NULL PRIMARY KEY,
	Key varchat(128),
	ValueStr TEXT,
	ValueInt BigInt,
	ValueFloat REAL,
	Date1 TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)
`

// interface wisejumper database table model;
type IWJDBTM_Helper1 interface {
	IWJDBTM_HelperCommon
	InsertIDKeyValue(id int64, key, value string) error
	SelectIDKeyValueTime() (*sql.Rows, error) // id, key, valuestr, date1
}
