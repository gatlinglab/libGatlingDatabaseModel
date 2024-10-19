package libGatlingDatabaseModel

import (
	"github.com/gatlinglab/libGatlingDatabaseModel/dbModel"
	idbModel "github.com/gatlinglab/libGatlingDatabaseModel/internal"
)

func GDM_CreateSqlDB(constr, token string) dbModel.IWJDatabase {
	return idbModel.NewDbModel(constr, token)
}

func TDM_CreateHelper1(dbInst dbModel.IWJDatabase, tablename string) dbModel.IWJDBTM_Helper1 {
	return idbModel.NewTableHelper1(dbInst, tablename)
}
