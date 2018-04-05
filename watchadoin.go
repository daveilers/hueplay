package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os/exec"
	"time"
)

type appInfo struct {
	Path     string
	BundleID string
	Name     string
}

func watchaDoin(a chan appInfo) {
	app := appInfo{}

	cmd := exec.Command(whatChaDoingCmd[0], whatChaDoingCmd[1:]...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		b := scanner.Bytes()
		if err := json.Unmarshal(b, &app); err != nil {
			log.Fatal(err)
		}
		a <- app
	}

	if err = cmd.Wait(); err != nil {
		log.Fatal(err)
	}

}

func (a *appInfo) isNotWork() bool {
	nw, ok := notWork[a.BundleID]
	return nw && ok
}
func (a *appInfo) isWork() bool {
	nw, ok := notWork[a.BundleID]
	return !nw && ok
}

func pollWhatchadDoin(events chan Event) {
	a := make(chan appInfo)
	go handleWhatchaDoin(events, a)
	for {
		watchaDoin(a)
		time.Sleep(appCheckPollRate)
	}
}

func handleWhatchaDoin(events chan Event, a chan appInfo) {
	last := appInfo{}
	for {
		app := <-a
		if app.BundleID != last.BundleID {
			last = app
			if app.isNotWork() {
				events <- NotWorking
				log.Printf("NotWorking: %+v", app)
				continue
			}
			if app.isWork() {
				events <- Working
				log.Printf("Working: %+v", app)
				continue
			}
			log.Printf("Schodinger's work:  %+v", app)
			events <- unknownWorking
		}
	}
}
