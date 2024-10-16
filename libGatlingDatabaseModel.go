package libGatlingDatabaseModel

import (
	"github.com/gatlinglab/libGatlingDatabaseModel/dbModel"
	idbModel "github.com/gatlinglab/libGatlingDatabaseModel/internal"
)

func GDM_CreateSqlDB(constr, token string) dbModel.IWJDatabase {
	return idbModel.NewDbModel(constr, token)
}
