package idbModel

import (
	"strings"

	"github.com/gatlinglab/libGatlingDatabaseModel/dbModel"
)

// type cDatabaseModelWrapper struct {
// }

func NewDbModel(constr, token string) dbModel.IWJDatabase {
	strTmp := strings.ToLower(constr)
	if strings.HasPrefix(strTmp, "libsql://") {
		return newDBModelTursoSql(constr, token)
	}

	return nil
}
