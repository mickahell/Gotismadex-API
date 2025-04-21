package docs

import (
	_ "embed"

	"github.com/mickahell/Gotismadex/logger"
)

//go:embed VERSION.txt
var version_file string

const draw = ``

func DrawStart() {
	logger.GetLogger().LogDraw(draw)
	logger.GetLogger().LogDraw("Version : " + version_file + "\n")
}

func GetVersion() string {
	return version_file
}
