package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
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

var logName = `hueplay.log`

func main() {
	f, err := os.OpenFile(logName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0700)
	if err != nil {
		log.Fatalf("I couldn't open the log file :-( :%v", err)
	}
	log.SetOutput(f)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		f.Close()
		os.Exit(0)
	}()

	events := make(chan Event)
	go checkBack(events, false)
	go pollWhatchadDoin(events)
	go sleepMonitor(events)
	handleEvents(events)
}
