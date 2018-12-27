package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/joho/godotenv/autoload"
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

var enableIdleCheck = flag.Bool("i", false, "if true dim on idle")
var enableLightControl = flag.Bool("l", false, "if true control lights")
var enableApplicationCheck = flag.Bool("a", false, "if true change color based on applications used")

func main() {
	flag.Parse()
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
	if *enableIdleCheck {
		go checkIdle(events, false)
	}
	if *enableApplicationCheck {
		go pollWhatchadDoin(events)
	}
	go sleepMonitor(events)
	handleEvents(events)
}
