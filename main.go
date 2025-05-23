package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/mickahell/Gotismadex/database"
	"github.com/mickahell/Gotismadex/docs"
	"github.com/mickahell/Gotismadex/helpers"
	"github.com/mickahell/Gotismadex/logger"
	"github.com/mickahell/Gotismadex/modules/launcher"
	"github.com/mickahell/Gotismadex/router"
)

func startoptions() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatalln("Abs path doesn't exist !")
	}

	flag.StringVar(
		&helpers.Confpathflag,
		"conf",
		filepath.Join(path, "conf/conf_local.yaml"),
		"path for the configuration file.",
	)
	flag.StringVar(
		&helpers.Swaggerpathflag,
		"swagger",
		"/conf/swagger.yaml",
		"relative path for the swagger file.",
	)

	flag.Parse()
}

func InitConf() {
	helpers.InitFile()
	helpers.ReadConfig()
	logger.GetLogger().LogInit()
}

func main() {
	docs.DrawStart()
	startoptions()
	InitConf()
	database.DatabaseInit()
	router.InitializeRouter()

	// Loading modules
	launcher.LauncherModules()

	logger.GetLogger().LogInfo("main", "Conf and modules loaded ; app starting...")

	// Specimen data
	if helpers.TheAppConfig().Specimen {
		launcher.SpecimenModules()
	}

	logger.GetLogger().LogInfo("main", "API ready.")

	// create a WaitGroup
	wg := new(sync.WaitGroup)
	// add two goroutines to `wg` WaitGroup
	wg.Add(2)

	go func() {
		logger.GetLogger().LogCritical("main", "listen error", router.GetServer().ListenAndServe())
		wg.Done()
	}()

	// wait until WaitGroup is done
	wg.Wait()
}
