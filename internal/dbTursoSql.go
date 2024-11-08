package idbModel

import (
	"database/sql"
	"errors"

	"github.com/gatlinglab/libGatlingDatabaseModel/dbModel"
	"github.com/tursodatabase/libsql-client-go/libsql"
)

type cDBModelTursoSql struct {
	connectStr   string
	connectToken string
	//dbConnector  *driver.Connector
	database         *sql.DB
	selfDatabaseType dbModel.DBMWJDatabaseType
	lastError        error
}

func newDBModelTursoSql(constr, token string) *cDBModelTursoSql {
	return &cDBModelTursoSql{connectStr: constr, connectToken: token, selfDatabaseType: dbModel.DBMWJDT_Sqlite, lastError: nil}
}

func (pInst *cDBModelTursoSql) Connect() error {
	connector, err := libsql.NewConnector(pInst.connectStr, libsql.WithAuthToken(pInst.connectToken))
	if err != nil {
		return err
	}
	//pInst.dbConnector = connector

	db := sql.OpenDB(connector)
	if db == nil {
		return errors.New("db open error")
	}
	pInst.database = db

	return nil
}
func (pInst *cDBModelTursoSql) Close() {
	//pInst.dbConnector.Close()
	pInst.database.Close()
}
func (pInst *cDBModelTursoSql) GetDatabaseType() dbModel.DBMWJDatabaseType {
	return dbModel.DBMWJDT_Sqlite
}

func (pInst *cDBModelTursoSql) ExecSql(sql string, args ...any) (sql.Result, error) {
	return pInst.database.Exec(sql, args)
}
func (pInst *cDBModelTursoSql) Query(sql string) (*sql.Rows, error) {
	return pInst.database.Query(sql)
}

func (pInst *cDBModelTursoSql) GetDatabaseVersion() (string, error) {
	pInst.lastError = nil
	rows, err := pInst.database.Query(dbHelperSqlCheckDatabaseVersion[int(dbModel.DBMWJDT_Sqlite)-1])
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

func (pInst *cDBModelTursoSql) CheckTableExists(tableName string) bool {
	pInst.lastError = nil
	rows, err := pInst.database.Query("SELECT name FROM sqlite_master WHERE type='table' AND name='" + tableName + "';")
	if err != nil {
		pInst.lastError = err
		return false
	}
	defer rows.Close()
	return rows.Next()
}
func (pInst *cDBModelTursoSql) DropTableIfExists(tableName string) error {
	strSql := "DROP TABLE IF EXISTS " + tableName
	_, err := pInst.database.Query(strSql)
	return err
}

/*
func (pInst *cDBModelTursoSql) GetLastError() error {
	return pInst.lastError
}
func (pInst *cDBModelTursoSql) CreateTableIfNotExists(tableType dbModel.DBMWJTableType, tableName string) error {
	strSql := getSingleTableHelperManager().TableHelperCreateTableSql(tableType, pInst.selfDatabaseType, tableName)
	if strSql == "" {
		return errors.New("table type error")
	}
	_, err := pInst.database.Exec(strSql)

	return err
}*/
