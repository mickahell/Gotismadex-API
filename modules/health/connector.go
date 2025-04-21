package health

import (
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/mickahell/Gotismadex/database"
	"github.com/mickahell/Gotismadex/docs"
	"github.com/mickahell/Gotismadex/helpers"
	"github.com/mickahell/Gotismadex/logger"
	"github.com/mickahell/Gotismadex/router"
)

func Launcher() {
	InitRoutes()
}

func TheLogger() *logger.Logger {
	return logger.GetLogger()
}

func TheRouter() *mux.Router {
	return router.GetRouter()
}

func TheSecureRouter() *mux.Router {
	return router.GetSecureRouter()
}

func TheAppConfig() *helpers.Cfg {
	return helpers.TheAppConfig()
}

func TheDb() *sql.DB {
	return database.Db()
}

func TheVersion() string {
	return docs.GetVersion()
}
