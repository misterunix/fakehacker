package ansicolor

import "fmt"

var BColor map[int]string
var FColor map[int]string

func Init() {
	BColor = make(map[int]string)
	FColor = make(map[int]string)

	for i := 0; i < 256; i++ {
		BColor[i] = fmt.Sprintf("\033[48;5;%dm", i)
		FColor[i] = fmt.Sprintf("\033[38;5;%dm", i)
	}

}
