package main

import (
	"time"
)

var notWork = map[string]bool{
	"com.apple.iChat":       true,
	"com.microsoft.VSCode":  false,
	"com.google.Chrome":     false,
	"com.apple.Terminal":    false,
	"com.microsoft.rdc.mac": false,
	"com.apple.dt.Xcode":    false,
	"com.sqlopsstudio.oss":  false,
}

const minIdleCheck = 5 * time.Second
const idleTimeOut = 2 * time.Minute
const sleepCheckInterval = 200 * time.Millisecond
const appCheckPollRate = time.Second

var whatChaDoingCmd = []string{"whatchaDoin"}

// const whatChaDoingPythoinCmd = []string{"python", "-u", "watcher.py"}

func main() {
	events := make(chan Event)
	go checkBack(events, false)
	go pollWhatchadDoin(events)
	go sleepMonitor(events)
	handleEvents(events)
}
