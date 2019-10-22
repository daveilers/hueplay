package main

import (
	"fmt"
	"os/exec"
)

var musicApps = []string{"iTunes", "Music", "Spotify"}

func pauseMusic() {
	p := exec.Command("osascript")
	for _, a := range musicApps {
		p.Args = append(p.Args,
			"-e", fmt.Sprintf(`if application %q is running then`, a),
			"-e", fmt.Sprintf(`tell application %q `, a),
			"-e", "pause",
			"-e", `end tell`,
			"-e", `end if`,
		)
	}
	p.Start()
}

func playMusic() {
	p := exec.Command("osascript")
	for _, a := range musicApps {
		p.Args = append(p.Args,
			"-e", fmt.Sprintf(`if application %q is running then`, a),
			"-e", fmt.Sprintf(`tell application %q `, a),
			"-e", "play",
			"-e", `end tell`,
			"-e", `end if`,
		)
	}
	p.Start()
	// exec.Command("osascript",
	// 	"-e", `if application "Music" is running then`,
	// 	"-e", `tell application "Music"`,
	// 	"-e", `play`,
	// 	"-e", `end tell`,
	// 	"-e", `end if`).Start()
}
