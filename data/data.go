package data

import "fmt"

// I hate globals

// Pop : type struct for holding popups
type Pop struct {
	Name     string
	Msg      string
	OnScreen bool
	Flashing bool
	X0, X1   int
	Y0, Y1   int
}

// PopUps : slice of all the popup windows
var PopUps []Pop

func Init() {

	CreatePops("ERROR", true)
	CreatePops("FAILURE", false)
	CreatePops("INTRUSION DETECTED", false)
	CreatePops("CRITICAL ERROR", false)
	CreatePops("ANTIVIRUS RUNNING", false)
	CreatePops("EXTERNAL ACCESS DETECTED", true)
	CreatePops("SHUTDOWN Y/n", false)
	CreatePops("VIRUS DETECTED", false)

}

func CreatePops(msg string, flashing bool) {
	p := Pop{}
	p.Msg = msg
	p.Name = fmt.Sprintf("pop%d", len(PopUps))
	p.OnScreen = false
	p.Flashing = flashing
	PopUps = append(PopUps, p)
}
