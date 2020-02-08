package blocks

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
)


func News(button string) {
	if err := exec.Command("pgrep","newsboat").Run(); err == nil { return }
	out, _ := exec.Command("newsboat","-x","print-unread").Output()
	re := regexp.MustCompile(`([[:digit:]]) .*\n`)
	if number, _ := strconv.Atoi(re.ReplaceAllString(string(out),"$1")); number > 0 {
		fmt.Printf(" %d\n",number)
	}

	if button == "1" {
		cmd := exec.Command("i3-msg","-q","exec $TERMINAL -e newsboat").Start()
		_ = cmd
	}
}
