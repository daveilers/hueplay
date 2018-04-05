package main

type Event int

const (
	Awake Event = iota
	Asleep
	Working
	NotWorking
	unknownWorking
	Active
	Idle
)

type DaveState struct {
	Awake   bool
	Active  bool
	Working bool
}

func handleEvents(events chan Event) {
	l := getLights()
	cl := color(l)
	turnEmOn(l)

	d := DaveState{Awake: true, Active: true, Working: true}
	for e := range events {
		switch e {
		case Awake:
			d.Awake = true
			playITunes()
			turnEmOn(l)
		case Asleep:
			d.Awake = false
			pauseITunes()
			turnEmOff(l)
		case Active:
			if d.Awake {
				d.Active = true
				active(l)
			}
		case Idle:
			if d.Awake {
				d.Active = false
				idle(l)
			}
		case Working:
			if d.Awake {
				d.Working = true
				working(cl)
			}
		case unknownWorking:
			break
		case NotWorking:
			if d.Awake {
				d.Working = false
				notWorking(cl)
			}
		}
	}
}
