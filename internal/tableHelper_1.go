package idbModel

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/gatlinglab/libGatlingDatabaseModel/dbModel"
)

var dbHelperSqlCreateTable1 = [int(dbModel.DBMWJDT_MAXINDEX) - 1]string{
	` CREATE TABLE IF NOT EXISTS %s(
		id                    BigInt         NOT NULL PRIMARY KEY,
		key varchar(128),
		valuestr TEXT,
		valueint BigInt,
		valuefloat REAL,
		date1 TIMESTAMP DEFAULT CURRENT_TIMESTAMP);`,
	` CREATE TABLE IF NOT EXISTS %s(
			id                    BigInt         NOT NULL PRIMARY KEY,
			key varchar(128),
			valuestr TEXT,
			valueint BigInt,
			valuefloat REAL,
			date1 TIMESTAMP DEFAULT NOW());`,
}

type CTableHelper1 struct {
	databaseType    dbModel.DBMWJDatabaseType
	dbInst          dbModel.IWJDatabase
	tableName       string
	helpDBDataIndex int
	insertCache     string
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
	tablename = strings.ToLower(strings.TrimSpace(tablename))
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
func (pInst *CTableHelper1) PutCacheIDKeyValue(id int64, key, value string) {
	pInst.insertCache = pInst.insertCache + "insert into " + pInst.tableName + "(id, key, valuestr) values (" +
		strconv.FormatInt(id, 10) + fmt.Sprintf(",'%s', '%s');\n", key, value)
}
func (pInst *CTableHelper1) ExecPutCache() error {
	_, err := pInst.dbInst.ExecSql(pInst.insertCache)
	if err == nil {
		pInst.insertCache = ""
	}

	return err
}

func (pInst *CTableHelper1) InsertIDKeyValue(id int64, key, value string) error {
	if pInst.dbInst == nil {
		return errors.New("no database instance")
	}

	strSql := "insert into " + pInst.tableName + "(id, key, valuestr) values (?,'?','?');"

	_, err := pInst.dbInst.ExecSql(strSql, id, key, value)

	return err
}
func (pInst *CTableHelper1) SelectIDKeyValueTime() (*sql.Rows, error) {
	sql := "select id, key, valuestr, date1 from " + pInst.tableName

	return pInst.dbInst.Query(sql)
}
