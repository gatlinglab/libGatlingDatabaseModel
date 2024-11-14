package libGatlingDatabaseModel

import (
	"fmt"
	"testing"
	"time"

	idbModel "github.com/gatlinglab/libGatlingDatabaseModel/internal"
)

// data@turso.serv00.net;
// const C_DBurl = "libsql://mydata-mydata.turso.io"
// const C_DBToken = "eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJhIjoicnciLCJpYXQiOjE3MjkwNTIzOTksImlkIjoiY2VlNGUyMGUtMTAxMS00Y2U3LTk2NDYtZmY4OTdlMzIwOGFmIn0.TkDYmlNLPMKLXpy1HM-SFaKnLBMATz1h8utd2mTTbQKdV82v6vhTqV0vr58w59TP2r3nXr62QQYPwhupYzIWDQ"

// account: nhost.io: data@turso.serv00.net;
const C_DBurl = "postgresql://796a357a-7cb7-4808-a627-9a836f760ef2-user:pw-5948c376-8cab-448f-b22a-9a0b362792c4@postgres-free-tier-v2020.gigalixir.com:5432/796a357a-7cb7-4808-a627-9a836f760ef2"

// "postgres://postgres:sn9JbUemd2YAvrTd@bsupuevsulhpmyulypgt.db.eu-central-1.nhost.run:5432/bsupuevsulhpmyulypgt"
const C_DBToken = ""

func TestGDM_ShowTableData(t *testing.T) {
	dbInst := GDM_CreateSqlDB(C_DBurl, C_DBToken)
	if dbInst == nil {
		t.Errorf("GDM_CreateSqlDB() error")
		return
	}

	err := dbInst.Connect(C_DBToken)
	if err != nil {
		t.Error("database connect failed: ", err)
		return
	}
	dbInst.SetTimeOutSeconds(30 * time.Second)

	const testTableName = "dailytest"
	tableHelp1 := idbModel.NewTableHelper1(dbInst, testTableName)
	if tableHelp1 == nil {
		t.Error("table helper create error.")
		return
	}

	fmt.Println("query 1")
	rows, err := tableHelp1.SelectIDKeyValueTime()
	fmt.Println("query 2")
	if err != nil {
		t.Error("select id, key, value error: ", err)
		return
	}

	iCount := 0
	for rows.Next() {
		iCount++
		var id int64
		var strKey string
		var strValue string
		var addTime time.Time
		rows.Scan(&id, &strKey, &strValue, &addTime)
		fmt.Println("loaded data from table: ", id, strKey, strValue, addTime)
	}
	fmt.Println("loaded data total rows: ", iCount)

}

func TestGDM_CreateSqlDB(t *testing.T) {
	dbInst := GDM_CreateSqlDB(C_DBurl, C_DBToken)
	if dbInst == nil {
		t.Errorf("GDM_CreateSqlDB() error")
		return
	}

	err := dbInst.Connect(C_DBToken)
	if err != nil {
		t.Error("database connect failed: ", err)
		return
	}

	version, err := dbInst.GetDatabaseVersion()
	if err != nil {
		t.Error("database get version failed: ", err)
		return
	}

	t.Logf("Version: %s\n", version)

	const testTableName = "testtable1"

	tableHelp1 := idbModel.NewTableHelper1(dbInst, testTableName)
	if tableHelp1 == nil {
		t.Error("table helper create error.")
		return
	}

	tableExists := tableHelp1.CheckTableExists()
	if tableExists {
		fmt.Println("table testTable1 exists")
	} else {
		fmt.Println("table testTable1 not exists")
	}

	if tableExists {
		err := tableHelp1.DropTableIfExists()
		if err != nil {
			t.Error(err)
			return
		} else {
			fmt.Println("table testTable1 exists, droped already;")
		}

		tableExists = tableHelp1.CheckTableExists()
		if tableExists {
			t.Error("table testTable1 exists after drop table")
			return
		} else {
			t.Log("table not exists now...")
		}
	}

	err = tableHelp1.CreateTable()
	if err != nil {
		t.Error("create table error", err)
		return
	} else {
		t.Log("table create successful")
	}

	tableExists = tableHelp1.CheckTableExists()
	if !tableExists {
		t.Error("table testTable1 not exists after create")
		return
	}

	//err = tableHelp1.InsertIDKeyValue(1, "testkey1", "testvalue1")
	err = tableHelp1.InsertKeyValue("testkey1", "testvalue1")
	//tableHelp1.PutCacheIDKeyValue(1, "testkey1", "testvalue1")
	//err = tableHelp1.ExecPutCache()
	if err != nil {
		t.Error("insert value error: ", err)
		return
	}

	rows, err := tableHelp1.SelectIDKeyValueTime()
	if err != nil {
		t.Error("select id, key, value error: ", err)
		return
	}

	iCount := 0
	for rows.Next() {
		iCount++
		if iCount != 1 {
			t.Error(" should be one row")
		}
		var id int64
		var strKey string
		var strValue string
		var addTime time.Time
		rows.Scan(&id, &strKey, &strValue, &addTime)
		fmt.Println("loaded data from table: ", id, strKey, strValue, addTime)
	}

	t.Log("successful\n")
}
