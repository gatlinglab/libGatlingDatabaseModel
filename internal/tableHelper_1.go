package idbModel

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"

	"github.com/gatlinglab/libGatlingDatabaseModel/dbModel"
)

var dbHelperSqlCreateTable1 = [int(dbModel.DBMWJDT_MAXINDEX) - 1]string{
	` CREATE TABLE IF NOT EXISTS %s(
		ID                    BigInt         NOT NULL PRIMARY KEY,
		Key varchat(128),
		ValueStr TEXT,
		ValueInt BigInt,
		ValueFloat REAL,
		Date1 TIMESTAMP DEFAULT CURRENT_TIMESTAMP)`,
}

type CTableHelper1 struct {
	databaseType    dbModel.DBMWJDatabaseType
	dbInst          dbModel.IWJDatabase
	tableName       string
	helpDBDataIndex int
}

func NewTableHelper1(db dbModel.IWJDatabase, tablename string) *CTableHelper1 {
	if db == nil {
		return nil
	}
	dbType := db.GetDatabaseType()
	if dbType == dbModel.DBMWJDT_Unknow || dbType == dbModel.DBMWJDT_MAXINDEX {
		return nil
	}
	if tablename == "" {
		return nil
	}
	return &CTableHelper1{dbInst: db, databaseType: dbType, tableName: tablename, helpDBDataIndex: int(dbType) - 1}
}

func (pInst *CTableHelper1) CreateTable() error {
	sql := fmt.Sprintf(dbHelperSqlCreateTable1[pInst.helpDBDataIndex], pInst.tableName)

	_, err := pInst.dbInst.ExecSql(sql)

	return err
}
func (pInst *CTableHelper1) CheckTableExists() bool {
	return pInst.dbInst.CheckTableExists(pInst.tableName)
}
func (pInst *CTableHelper1) DropTableIfExists() error {
	return pInst.dbInst.DropTableIfExists(pInst.tableName)
}

func (pInst *CTableHelper1) InsertIDKeyValue(id int64, key, value string) error {
	if pInst.dbInst == nil {
		return errors.New("no database instance")
	}
	strSql := "insert into " + pInst.tableName + "(ID, Key, ValueStr) values (" + strconv.FormatInt(id, 10)
	strSql += fmt.Sprintf(",\"%s\", \"%s\");", key, value)

	_, err := pInst.dbInst.ExecSql(strSql)

	return err
}
func (pInst *CTableHelper1) SelectIDKeyValueTime() (*sql.Rows, error) {
	sql := "select ID, Key, ValueStr, Date1 from " + pInst.tableName

	return pInst.dbInst.Query(sql)
}
