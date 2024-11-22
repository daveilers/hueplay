package main

import (
	"log"
	"time"
)

/*
#include <stdbool.h>
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L. -lasleep
#include <asleep.h>
*/
import "C"

func asleep() bool {
	isAsleep := C.Asleep()
	return isAsleep == true
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
