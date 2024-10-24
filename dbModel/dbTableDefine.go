package dbModel

//import "time"

// id primary, Key string, ValueStr string, ValueInt int, ValueFloat float64,date default now;
// type CWJDBTableIDKVD struct {
// 	ID         int64
// 	Key        string
// 	ValueStr   string
// 	ValueInt   int64
// 	ValueFloat float64
// 	Date1      time.Time
// }

type DBMWJTableType int

const (
	DBMWJTT_Unknow DBMWJTableType = iota
	DBMWJTT_KeyAll1

	DBMWJTT_MAXINDEX
)

type DBMWJDatabaseType int

const (
	DBMWJDT_Unknow DBMWJDatabaseType = iota
	DBMWJDT_Sqlite
	DBMWJDT_Postgres

	DBMWJDT_MAXINDEX
)

type IWJDBTM_HelperCommon interface {
	CreateTable() error
	CheckTableExists() bool
	DropTableIfExists() error
	///// put insert sql text to cache; exec once by next function;
	PutCacheIDKeyValue(id int64, key, value string)
	ExecPutCache()error
	//// above;
}
