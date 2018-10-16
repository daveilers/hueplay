package main

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strconv"
	"time"
)

func checkIdle(c chan Event, wasIdle bool) {
	i, err := idleTime()
	if err != nil {
		log.Printf("Dave broke the idle checker: %v", err)
	}
	if i > idleTimeOut {
		if !wasIdle {
			log.Printf("Idle: You've been idle for %v", i)
			c <- Idle
		}
		<-time.After(minIdleCheck)
		checkIdle(c, true)
	} else {
		if wasIdle {
			log.Printf("Active: You've been idle for %v", i)
			c <- Active
		}
		<-time.After(idleTimeOut - i)
		checkIdle(c, false)
	}
}

func idleTime() (time.Duration, error) {
	idleThings := exec.Command("ioreg", "-c", "IOHIDSystem")
	d, err := idleThings.Output()
	if err != nil {
		return 0, err
	}
	idleLine := regexp.MustCompile(`"HIDIdleTime" = ([0-9]+)`)
	m := idleLine.FindSubmatch(d)
	if len(m) < 2 {
		return 0, fmt.Errorf("I can't find the value I'm looking for")
	}
	idleNano, err := strconv.Atoi(string(m[1]))
	if err != nil {
		return 0, err
	}
	return time.Duration(idleNano), nil
}
