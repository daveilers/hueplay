package main

import (
	"log"
	"os/exec"
	"time"
)

func asleep() bool {
	// asleep := exec.Command(`python`, `-c`, `import sys,Quartz; d=Quartz.CGSessionCopyCurrentDictionary(); sys.exit(d and d.get("CGSSessionScreenIsLocked", 0) == 0 and d.get("kCGSSessionOnConsoleKey", 0) == 1)`)
	asleep := exec.Command(`asleep`)
	err := asleep.Run()
	return err == nil
}

func sleepMonitor(events chan Event) {
	c := time.Tick(sleepCheckInterval)
	wasSleeping := false
	for now := range c {
		if asleep() {
			if !wasSleeping {
				log.Printf("Noticed I was asleep at %v", now)
				events <- Asleep
				wasSleeping = true
			}
		} else {
			if wasSleeping {
				log.Printf("Noticed I was awake at %v", now)
				events <- Awake
				wasSleeping = false
			}
		}
	}

}
