package idbModel

import (
	"database/sql"
	"fmt"
	"net/url"

	"github.com/gatlinglab/libGatlingDatabaseModel/dbModel"
	_ "github.com/lib/pq"
)

type cDBModelPostgres struct {
	connectStr       string
	database         *sql.DB
	selfDatabaseType dbModel.DBMWJDatabaseType
	lastError        error
}

func newDBModelPostgres(constr string) *cDBModelPostgres {
	return &cDBModelPostgres{connectStr: constr, selfDatabaseType: dbModel.DBMWJDT_Postgres, lastError: nil}
}

func (pInst *cDBModelPostgres) Connect() error {
	conn, _ := url.Parse(pInst.connectStr)
	conn.RawQuery = "sslmode=disable"

	db, err := sql.Open("postgres", conn.String())
	if err != nil {
		return err
	}

	pInst.database = db

	return nil
}
func (pInst *cDBModelPostgres) Close() {
	//pInst.dbConnector.Close()
	pInst.database.Close()
}
func (pInst *cDBModelPostgres) GetDatabaseType() dbModel.DBMWJDatabaseType {
	return dbModel.DBMWJDT_Postgres
}

func (pInst *cDBModelPostgres) ExecSql(sql string, args ...any) (sql.Result, error) {
	return pInst.database.Exec(sql, args)
}
func (pInst *cDBModelPostgres) Query(sql string) (*sql.Rows, error) {
	return pInst.database.Query(sql)
}

func (pInst *cDBModelPostgres) GetDatabaseVersion() (string, error) {
	pInst.lastError = nil
	rows, err := pInst.database.Query(dbHelperSqlCheckDatabaseVersion[int(dbModel.DBMWJDT_Postgres)-1])
	if err != nil {
		pInst.lastError = err
		return "", err
	}

	var result string
	for rows.Next() {
		err = rows.Scan(&result)
		if err != nil {
			pInst.lastError = err
			return "", err
		}
	}
	return result, nil
}

func (pInst *cDBModelPostgres) CheckTableExists(tableName string) bool {
	pInst.lastError = nil
	// query := `SELECT EXISTS (
	// 	SELECT 1
	// 	FROM   information_schema.tables
	// 	WHERE  table_name = $1
	// );`
	//strSql := query + tableName + "');"
	var exists bool
	//query := fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM pg_catalog.pg_class c JOIN pg_catalog.pg_namespace n ON n.oid = c.relnamespace WHERE n.nspname = 'public' AND c.relname = '%s' AND c.relkind = 'r')", tableName)

	query := fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = '%s')", tableName)
	err := pInst.database.QueryRow(query).Scan(&exists)

	//var exists bool
	//strSql := "select * from information_schema.tables where table_name ='" + tableName + "';"
	//err := pInst.database.QueryRow(query, tableName).Scan(&exists)
	//fmt.Println(query)
	if err != nil {
		fmt.Println("check table exists error: ", err)
		pInst.lastError = err
		return false
	}

	return exists
}
func (pInst *cDBModelPostgres) DropTableIfExists(tableName string) error {
	strSql := "DROP TABLE IF EXISTS " + tableName
	_, err := pInst.database.Exec(strSql)
	return err
}
