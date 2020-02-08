package blocks

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)


func Nmblocks(button string) {
	typeRe:= regexp.MustCompile(`\w+ +(\w+).*`)
	isConnectedRe:= regexp.MustCompile(`.* connected .*`)
	out, _ := exec.Command("nmcli","device").Output()
	device := strings.Split(string(out),"\n")
	if isConnectedRe.MatchString(device[1]) {
		switch connectionType := typeRe.ReplaceAllString(device[1],"$1"); connectionType {
		case "ethernet":
			fmt.Println("")
		case "wifi":
			fmt.Println("")
		}
	} else {
		fmt.Println("")
	}

	if button == "1" {
		cmd := exec.Command("i3-msg","-q","exec $TERMINAL -e nmtui").Start()
		_ = cmd
	}
}
