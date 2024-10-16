package libGatlingDatabaseModel

import (
	"fmt"
	"testing"
)

func TestGDM_CreateSqlDB(t *testing.T) {
	dbInst := GDM_CreateSqlDB("libsql://mydata-mydata.turso.io", "eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJhIjoicnciLCJpYXQiOjE3MjkwNTIzOTksImlkIjoiY2VlNGUyMGUtMTAxMS00Y2U3LTk2NDYtZmY4OTdlMzIwOGFmIn0.TkDYmlNLPMKLXpy1HM-SFaKnLBMATz1h8utd2mTTbQKdV82v6vhTqV0vr58w59TP2r3nXr62QQYPwhupYzIWDQ")
	if dbInst == nil {
		t.Errorf("GDM_CreateSqlDB() error")
		return
	}

	err := dbInst.Connect()
	if err != nil {
		t.Error("database connect failed: ", err)
		return
	}

	rows, err := dbInst.Query("SELECT sqlite_version();") //"select version()")
	if err != nil {
		t.Error("SELECT version() error", err)
		return
	}

	for rows.Next() {
		var result string
		err = rows.Scan(&result)
		if err != nil {
			t.Error("rows scan error in SELECT version(): ", err)
			return
		}
		fmt.Printf("Version: %s\n", result)
		t.Logf("Version: %s\n", result)
	}

	t.Log("successful\n")
}
