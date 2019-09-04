package main

import (
	"flag"
	"fmt"
	"github.com/Margular/L4D2-VoicePack-Installer/steam/l4d2"
)

func main() {
	var steamPath = flag.String("path", "", "steam installation path")

	flag.Parse()

	var s = l4d2.NewL4d2()
	s.SetSteamPath(*steamPath)
	var err = s.InstallVoicePack()
	if err != nil {
		panic(err)
	}

	fmt.Println("voice pack install successfully! now type snd_rebuildaudiocache in the left for dead 2!")
}
