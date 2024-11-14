package idbModel

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"
	"time"

	"github.com/gatlinglab/libGatlingDatabaseModel/dbModel"
	_ "github.com/lib/pq"
)

type cDBModelPostgres struct {
	connectStr       string
	connectToken     string
	database         *sql.DB
	selfDatabaseType dbModel.DBMWJDatabaseType
	lastError        error
	timeoutSecond    time.Duration
}

func newDBModelPostgres(constr, conToken string) *cDBModelPostgres {
	return &cDBModelPostgres{connectStr: constr, selfDatabaseType: dbModel.DBMWJDT_Postgres, lastError: nil, timeoutSecond: 10 * time.Second}
}

func (pInst *cDBModelPostgres) Connect(sslConfig string) error {
	conn, _ := url.Parse(pInst.connectStr)
	//conn.RawQuery = "sslmode=disable"
	conn.RawQuery = sslConfig

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
func (pInst *cDBModelPostgres) SetTimeOutSeconds(timeout time.Duration) {
	pInst.timeoutSecond = timeout
}
func (pInst *cDBModelPostgres) GetDatabaseType() dbModel.DBMWJDatabaseType {
	return dbModel.DBMWJDT_Postgres
}
func (pInst *cDBModelPostgres) GetDBHandler() *sql.DB {
	return pInst.database
}

func (pInst *cDBModelPostgres) ExecSql(sql string, args ...any) (sql.Result, error) {
	context, _ := context.WithTimeout(context.Background(), pInst.timeoutSecond)
	return pInst.database.ExecContext(context, sql, args...)
}
func (pInst *cDBModelPostgres) Query(sql string) (*sql.Rows, error) {
	context, _ := context.WithTimeout(context.Background(), pInst.timeoutSecond)
	rows, err := pInst.database.QueryContext(context, sql)
	return rows, err
}

func (pInst *cDBModelPostgres) GetDatabaseVersion() (string, error) {
	pInst.lastError = nil
	context, _ := context.WithTimeout(context.Background(), pInst.timeoutSecond)
	rows, err := pInst.database.QueryContext(context, dbHelperSqlCheckDatabaseVersion[int(dbModel.DBMWJDT_Postgres)-1])
	if err != nil {
		pInst.lastError = err
		return "", err
	}
	defer rows.Close()

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

	context, _ := context.WithTimeout(context.Background(), pInst.timeoutSecond)
	query := fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = '%s')", tableName)
	err := pInst.database.QueryRowContext(context, query).Scan(&exists)

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
	context, _ := context.WithTimeout(context.Background(), pInst.timeoutSecond)
	_, err := pInst.database.ExecContext(context, strSql)
	return err
}
