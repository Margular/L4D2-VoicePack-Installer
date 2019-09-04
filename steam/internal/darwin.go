// +build darwin

package internal

import "os/user"

func init() {
	var u, err = user.Current()
	if err != nil {
		panic(err)
	}

	SteamPath = "/Users/" + u.Username + "/Library/Application Support/Steam"
}

var SteamPath string
