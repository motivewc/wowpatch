package platform

import (
	"path/filepath"
	"runtime"
)

func FindWarcraftClientExecutable() string {
	switch runtime.GOOS {
	case "windows":
		return "" // todo: idk
	case "darwin":
		return filepath.Join(
			"/",
			"Applications",
			"World of Warcraft",
			"_retail_",
			"World of Warcraft.app",
			"Contents",
			"MacOS",
		)
	}
	return ""
}
