package main

import (
	"os/exec"
)

var musicApps = []string{"iTunes", "Music", "Spotify"}

func pauseITunes() {
	exec.Command("osascript",
		"-e", `if application "Music" is running then`,
		"-e", `tell application "Music"`,
		"-e", "pause",
		"-e", `end tell`,
		"-e", `end if`).Start()
}

func playITunes() {
	exec.Command("osascript",
		"-e", `if application "Music" is running then`,
		"-e", `tell application "Music"`,
		"-e", `play`,
		"-e", `end tell`,
		"-e", `end if`).Start()
}
