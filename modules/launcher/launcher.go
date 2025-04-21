package launcher

import (
	"github.com/mickahell/Gotismadex/logger"
	"github.com/mickahell/Gotismadex/modules/health"
)

// Modules to launch, the order can be important!
func LauncherModules() {

	// Always last
	health.Launcher()
}

// Populate database, the order can be important!
func SpecimenModules() {

	logger.GetLogger().LogInfo("Specimen", "Specimen data charged up.")
}
