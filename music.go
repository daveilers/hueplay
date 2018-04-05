package main

import (
	"os/exec"
)

func pauseITunes() {
	exec.Command("osascript",
		"-e", `if application "iTunes" is running then`,
		"-e", `tell application "iTunes"`,
		"-e", "pause",
		"-e", `end tell`,
		"-e", `end if`).Start()
}

func playITunes() {
	exec.Command("osascript",
		"-e", `if application "iTunes" is running then`,
		"-e", `tell application "iTunes"`,
		"-e", `play`,
		"-e", `end tell`,
		"-e", `end if`).Start()
}
