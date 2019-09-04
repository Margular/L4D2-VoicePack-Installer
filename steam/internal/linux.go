// +build linux

package internal

import "os/user"

func init() {
	var u, err = user.Current()
	if err != nil {
		panic(err)
	}

	if u.Username == "root" {
		SteamPath = "/root/.steam/steam"
	} else {
		SteamPath = "/home/" + u.Username + "/.steam/steam"
	}
}

var SteamPath string
