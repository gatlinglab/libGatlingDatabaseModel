package idbModel

import (
	"github.com/gatlinglab/libGatlingDatabaseModel/dbModel"
)

var dbHelperSqlCheckDatabaseVersion = [int(dbModel.DBMWJDT_MAXINDEX) - 1]string{
	"SELECT sqlite_version()",
	"SELECT version()",
}

type ITableHelper interface {
	MakeTableCreateSql(dbType dbModel.DBMWJDatabaseType, tableName string) string
}

/*
func init() {
	//g_TableHelperList[0] = nil
	//g_TableHelperList[1] = newTableHelper1()
	getSingleTableHelperManager().Initialize()
}

type CTableHelperManager struct {
	tHelperList [int(dbModel.DBMWJTT_MAXINDEX)]ITableHelper
}

var g_singleTableHelperManager *CTableHelperManager = &CTableHelperManager{}

func getSingleTableHelperManager() *CTableHelperManager {
	return g_singleTableHelperManager
}

func (pInst *CTableHelperManager) Initialize() {
	pInst.tHelperList[0] = &cTableHelperEmpty{}
	pInst.tHelperList[1] = NewTableHelper1(nil, dbModel.DBMWJDT_Unknow, "")
}
func (pInst *CTableHelperManager) GetTableHelper(tableType dbModel.DBMWJTableType) ITableHelper {
	return pInst.tHelperList[int(tableType)]
}
func (pInst *CTableHelperManager) TableHelperCreateTableSql(tableType dbModel.DBMWJTableType, dbType dbModel.DBMWJDatabaseType, tableName string) string {
	return pInst.tHelperList[int(tableType)].MakeTableCreateSql(dbType, tableName)
}

type cTableHelperEmpty struct {
}

func (pInst *cTableHelperEmpty) MakeTableCreateSql(dbType dbModel.DBMWJDatabaseType, tableName string) string {
	return ""
}*/
