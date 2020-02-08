package blocks

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)


func Root(button string) {
	sizeRe:= regexp.MustCompile(`.* (\w+) +[[:digit:]]+% +/$`)
	out, _ := exec.Command("df","-h","-P","/").Output()
	lines := strings.Split(string(out),"\n")
	size := sizeRe.ReplaceAllString(lines[1],"$1")
	fmt.Printf(" %s\n",size)

	if button == "1" {
		cmd := exec.Command("i3-msg","-q","exec $TERMINAL -e ranger").Start()
		_ = cmd
	}
}
