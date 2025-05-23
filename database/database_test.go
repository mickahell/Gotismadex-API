package database

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/mickahell/Gotismadex/helpers"
	"github.com/mickahell/Gotismadex/logger"
)

func TestMain(m *testing.M) {
	setUp()
	//log.Println("Do stuff BEFORE the tests!")
	exitVal := m.Run()
	//log.Println("Do stuff AFTER the tests!")
	//tearDown()
	os.Exit(exitVal)
}

func setUp() {
	helpers.Confpathflag = "../conf/conf_local.yaml"
	helpers.InitFile()
	helpers.ReadConfig()
	logger.GetLogger().LogInit()
}

func TestDatabaseInit(t *testing.T) {
	DatabaseInit()
}

func TestDb(t *testing.T) {
	var got interface{} = Db()

	_, ok := got.(*sql.DB)

	if ok == false {
		log.Fatalln("got is not what i want !")
	}

}
