package blocks

import (
	"fmt"
	"os/exec"
	"os"
)


func Screenawake(button string) {
	if _, err := os.Stat("/tmp/lockscreen.lock"); ! os.IsNotExist(err) {
		fmt.Println("")
	}


	if (button == "1") || (button == "3") {
		if err := os.Remove("/tmp/lockscreen.lock"); err == nil {
			cmd := exec.Command("notify-send"," Automatic screenlock is now enabled").Run()
			_ = cmd
		}
	}

}
