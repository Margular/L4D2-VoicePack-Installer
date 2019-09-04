package steam

import (
	"fmt"
	"github.com/Margular/L4D2-VoicePack-Installer/steam/internal"
	"os"
)

type Steam struct {
	path string
}

func NewSteam() *Steam {
	return new(Steam)
}

func (s Steam) SetPath(path string) {
	s.path = path
}

func (s Steam) GetPath() string {
	if len(s.path) > 0 {
		fmt.Println("get cached steam path: " + s.path)
		return s.path
	}

	// Find steam path
	fmt.Println("now try to find steam path")
	s.path = s.findPath()
	if len(s.path) > 0 {
		fmt.Println("steam path found: " + s.path)
	} else {
		fmt.Println("steam path not found!")
	}
	return s.path
}

func (s Steam) findPath() string {
	var guessPath = internal.SteamPath

	if _, err := os.Stat(guessPath); !os.IsNotExist(err) {
		fmt.Println("steam installed in the default directory \"" + guessPath + "\"")
		return guessPath
	}

	fmt.Println("steam is not installed in the default directory")
	return ""
}
