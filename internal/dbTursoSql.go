package idbModel

import (
	"database/sql"
	"errors"

	"github.com/tursodatabase/libsql-client-go/libsql"
)

type cDBModelTursoSql struct {
	connectStr   string
	connectToken string
	//dbConnector  *driver.Connector
	database *sql.DB
}

func newDBModelTursoSql(constr, token string) *cDBModelTursoSql {
	return &cDBModelTursoSql{connectStr: constr, connectToken: token}
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

func (pInst *cDBModelTursoSql) ExecSql(sql string) (sql.Result, error) {
	return pInst.database.Exec(sql)
}
func (pInst *cDBModelTursoSql) Query(sql string) (*sql.Rows, error) {
	return pInst.database.Query(sql)
}
