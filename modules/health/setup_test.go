package health

import (
	"log"
	"os"
	"testing"

	"github.com/mickahell/Gotismadex/database"
	"github.com/mickahell/Gotismadex/helpers"
)

func TestMain(m *testing.M) {
	testMainSetup()

	//log.Println("Do stuff BEFORE the tests!")
	exitVal := m.Run()
	//log.Println("Do stuff AFTER the tests!")
	testMainTeardown()
	os.Exit(exitVal)

}

func testMainSetup() {
	helpers.Confpathflag = "../../conf/conf_local.yaml"
	helpers.InitFile()
	helpers.ReadConfig()
	TheLogger().LogInit()
	database.DatabaseInit()
}

func testMainTeardown() {
	_, err := TheDb().Exec("DROP DATABASE IF EXISTS " + helpers.TheAppConfig().Namedb + ";")
	if err != nil {
		log.Fatal(err)
	}
}
