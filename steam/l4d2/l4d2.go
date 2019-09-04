package l4d2

import (
	"fmt"
	"github.com/Margular/L4D2-VoicePack-Installer/steam"
	"github.com/Margular/vpk"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

type l4d2 struct {
	path  string
	steam steam.Steam
}

func NewL4d2() *l4d2 {
	return new(l4d2)
}

func (l l4d2) SetSteamPath(path string) {
	l.steam.SetPath(path)
}

func (l l4d2) GetPath() string {
	var steamPath = l.steam.GetPath()

	if steamPath == "" {
		return ""
	}

	l.path = steamPath + "/steamapps/common/Left 4 Dead 2"
	return l.path
}

func (l l4d2) InstallVoicePack() error {
	if len(l.path) == 0 {
		l.path = l.GetPath()

		if len(l.path) == 0 {
			return fmt.Errorf("can not find the directory path of left for dead 2")
		}
	}

	var addonsPath = l.path + "/left4dead2/addons"

	err := filepath.Walk(addonsPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if filepath.Ext(path) == ".vpk" {
				var v, err = vpk.Open(vpk.SingleVPK(path))

				if err != nil {
					fmt.Printf("error occurred when open %s %s\n", path, err)
					return nil
				}

				fmt.Println("now parsing " + path)

				for _, entryPath := range v.Paths() {
					matched, err := regexp.Match(`left4dead2(_dlc[123])?/sound/player/survivor/voice/.*`,
						[]byte(entryPath))
					if err != nil {
						fmt.Printf("error occurred when regexp on %s %s\n", entryPath, err)
						continue
					}

					if matched {
						fmt.Println("entry found: " + entryPath)

						entryFile, err := v.Open(entryPath)
						if err != nil {
							fmt.Printf("error occurred when open entry %s %s\n", entryPath, err)
							continue
						}

						data, err := ioutil.ReadAll(entryFile)
						if err != nil {
							fmt.Printf("error occurred when read data from %s %s\n", entryPath, err)
							continue
						}

						var targetPath = l.path + "/" + entryPath
						err = ioutil.WriteFile(targetPath, data, 0644)
						if err != nil {
							fmt.Printf("failed to write to file: %s %s\n", targetPath, err)
							continue
						}

						fmt.Printf("write to file %s successfully\n", targetPath)
					}
				}
			}

			return nil
		})
	if err != nil {
		return err
	}

	return nil
}
