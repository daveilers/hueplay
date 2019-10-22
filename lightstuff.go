package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/kelseyhightower/envconfig"

	"github.com/gbbr/hue"
)

// func workingHours(t time.Time) bool {
// 	day := t.Weekday()
// 	hour := t.Hour()
// 	fmt.Printf("day:%v hour:%v zone:%v\n", day, hour, t.Location())
// 	if hour >= 8 && hour < 17 && day > time.Sunday && day < time.Saturday {
// 		return true
// 	}
// 	return false
// }

func getBridge() *hue.Bridge {
	b, err := hue.Discover()
	if err != nil {
		log.Fatalln(err)
	}
	// log.Printf("%+v\n", b)
	return b
}

func working(l []*hue.Light) error {
	newState := hue.State{On: true, Brightness: uint8(254), Hue: uint16(65535), Saturation: uint8(254)}
	for _, light := range l {
		err := light.Set(&newState)
		if err != nil {
			return err
		}
	}
	return nil
}
func notWorking(l []*hue.Light) error {
	newState := hue.State{On: true, Brightness: uint8(254), Hue: uint16(25500), Saturation: uint8(254)}
	for _, light := range l {
		err := light.Set(&newState)
		if err != nil {
			return err
		}
	}
	return nil
}
func idle(l []*hue.Light) error {
	newState := hue.State{On: true,
		Brightness: 1,
		// Saturation: uint8(1),
		TransitionTime: uint16(5 * time.Minute / time.Millisecond / 100)}
	for _, light := range l {
		err := light.Set(&newState)
		if err != nil {
			return err
		}
	}
	return nil
}
func active(l []*hue.Light) error {
	newState := hue.State{
		On:         true,
		Brightness: 254,
		// Saturation: uint8(254), HueInc: 1
	}
	for _, light := range l {
		err := light.Set(&newState)
		if err != nil {
			return err
		}
	}
	return nil
}

func doSomething(l *hue.Light) {
	// colors := []*[2]float32{hue.GREEN, hue.RED, hue.YELLOW}

	newState := hue.State{On: true, Brightness: uint8(rand.Intn(253)) + 1, Hue: uint16(rand.Intn(65000)), Saturation: uint8(rand.Intn(253) + 1)}
	err := l.Set(&newState)
	if err != nil {
		log.Fatalln(err)
	}

}

func color(l []*hue.Light) []*hue.Light {
	cl := []*hue.Light{}
	for _, light := range l {
		if light.State.ColorMode != "" {
			cl = append(cl, light)
		}
	}
	return cl
}

func getLights() []*hue.Light {
	b := getBridge()
	ls := b.Lights()

	l, err := ls.List()
	if err != nil {
		log.Fatalln(err)
	}
	l = ignoreLights(l)
	return l
}

func ignoreLights(lights []*hue.Light) []*hue.Light {
	var igs = &ignorecfg{}
	envconfig.MustProcess("", igs)
	// fmt.Printf("%v\n", igs)

	keepLights := []*hue.Light{}
All:
	for i, light := range lights {
		for _, UIDToIgnore := range igs.Ignore {
			if light.UID == UIDToIgnore {
				continue All
			}
		}
		keepLights = append(keepLights, lights[i])
	}
	return keepLights
}

type ignorecfg struct {
	Ignore []string
}

func listLights(ll []*hue.Light) {
	for i, l := range ll {
		log.Printf("%v, %+v\n", i, l)
	}
}

func turnEmOn(l []*hue.Light) {
	for _, light := range l {
		light.On()
		newState := &hue.State{
			On:         true,
			Brightness: 254,
			// Saturation: uint8(254), HueInc: 1
		}
		light.Set(newState)
	}
}

func turnEmOff(l []*hue.Light) {
	for _, light := range l {
		light.Off()
	}
}
