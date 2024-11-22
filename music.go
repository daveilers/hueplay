package main

import (
	"github.com/andybrewer/mack"
)

var musicApps = []string{"Music", "Spotify"}

func pauseMusic() {
	for _, musicApp := range musicApps {
		_, _ = mack.Tell(musicApp, `pause`)
	}
}

func playMusic() {
	for _, musicApp := range musicApps {
		_, _ = mack.Tell(musicApp, `play`)
	}
}
