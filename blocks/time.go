package blocks

import (
	"fmt"
	"time"
	"os/exec"
)


func Time(button string) {
	t := time.Now()
	date := t.Format(" 02-01-2006  15:04")
	fmt.Println(date)

	switch button {
	case "1":
		cmd := exec.Command("i3-msg","-q","exec $TERMINAL -e calcurse && savetodownloadsongs.sh").Start()
		_ = cmd
	case "3":
		cmd := exec.Command("i3-msg","-q","exec $TERMINAL -e tty-clock -s -c -b").Start()
		_ = cmd
	}
}
