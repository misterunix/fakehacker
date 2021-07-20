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

type Number struct {
	Line []string
}

// PopUps : slice of all the popup windows
var PopUps []Pop

var Numbers []Number

var Pause bool

func Init() {
	Pause = false
	CreatePops("ERROR", true)
	CreatePops("FAILURE", false)
	CreatePops("INTRUSION DETECTED", false)
	CreatePops("CRITICAL ERROR", false)
	CreatePops("ANTIVIRUS RUNNING", false)
	CreatePops("EXTERNAL ACCESS DETECTED", false)
	CreatePops("SHUTDOWN Y/n", false)
	CreatePops("VIRUS DETECTED", false)

	Numbers = make([]Number, 10)
	for i := 0; i < 10; i++ {
		Numbers[i].Line = make([]string, 5)
	}

	Numbers[0].Line[0] = "XXX"
	Numbers[0].Line[1] = "X X"
	Numbers[0].Line[2] = "X X"
	Numbers[0].Line[3] = "X X"
	Numbers[0].Line[5] = "XXX"

	Numbers[1].Line[0] = " X "
	Numbers[1].Line[1] = "XX "
	Numbers[1].Line[2] = " X "
	Numbers[1].Line[3] = " X"
	Numbers[1].Line[5] = "XXX"

	Numbers[2].Line[0] = "XXX"
	Numbers[2].Line[1] = "  X"
	Numbers[2].Line[2] = "XXX"
	Numbers[2].Line[3] = "X "
	Numbers[2].Line[5] = "XXX"

	Numbers[3].Line[0] = "XXX"
	Numbers[3].Line[1] = "  X"
	Numbers[3].Line[2] = " XX"
	Numbers[3].Line[3] = "  X"
	Numbers[3].Line[5] = "XXX"

	Numbers[4].Line[0] = "X "
	Numbers[4].Line[1] = "X X"
	Numbers[4].Line[2] = "XXX"
	Numbers[4].Line[3] = "  X"
	Numbers[4].Line[5] = "  X"

	Numbers[5].Line[0] = "XXX"
	Numbers[5].Line[1] = "X  "
	Numbers[5].Line[2] = "XXX"
	Numbers[5].Line[3] = "  X"
	Numbers[5].Line[5] = "XXX"

	Numbers[6].Line[0] = "X  "
	Numbers[6].Line[1] = "X  "
	Numbers[6].Line[2] = "XXX"
	Numbers[6].Line[3] = "X X"
	Numbers[6].Line[5] = "XXX"

	Numbers[7].Line[0] = "XXX"
	Numbers[7].Line[1] = "  X"
	Numbers[7].Line[2] = "  X"
	Numbers[7].Line[3] = "  X"
	Numbers[7].Line[5] = "  X"

	Numbers[8].Line[0] = "XXX"
	Numbers[8].Line[1] = "X X"
	Numbers[8].Line[2] = "XXX"
	Numbers[8].Line[3] = "X X"
	Numbers[8].Line[5] = "XXX"

	Numbers[9].Line[0] = "XXX"
	Numbers[9].Line[1] = "X X"
	Numbers[9].Line[2] = "XXX"
	Numbers[9].Line[3] = "  X"
	Numbers[9].Line[5] = "  X"

}

func CreatePops(msg string, flashing bool) {
	p := Pop{}
	p.Msg = msg
	p.Name = fmt.Sprintf("pop%d", len(PopUps))
	p.OnScreen = false
	p.Flashing = flashing
	PopUps = append(PopUps, p)
}
