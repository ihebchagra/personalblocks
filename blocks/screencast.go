package blocks

import (
	"fmt"
	"os/exec"
)


func Screencast(button string) {
	switch button {
	case "1", "3":
		if err := exec.Command("pkill","-f","ffmpeg -f x11grab").Run(); err == nil {
			cmd := exec.Command("notify-send"," Killed Screencast").Run()
			_ = cmd
		}
	case "":
		if err := exec.Command("pgrep","-f","ffmpeg -f x11grab").Run(); err == nil {
			fmt.Println("")
		}
	}

}
