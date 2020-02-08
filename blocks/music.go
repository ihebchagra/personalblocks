package blocks

import (
	"fmt"
	"os/exec"
	"regexp"
	"time"
)


func Music(button string) {
	s, _ := exec.Command("mpc","current","-f","%file%").Output()
	file := string(s)
	re := regexp.MustCompile(`^https:\/\/lainon\.life\/radio\/.*\.ogg\n$`)
	if (file == "https://listen.moe/stream\n") || (file == "http://stream.pianoramaradio.ru\n") || re.MatchString(file) {
		return
	}
	format:="[%albumartist% - ]|[%artist% - ]&%title%"
	out, _:= exec.Command("mpc","-f",format,"current").Output()
	if song:=string(out); song != "" {
		fmt.Printf("%s",song)
	}

	switch button {
	case "1":
		cmd := exec.Command("i3-msg","-q","exec $TERMINAL -e ncmpcpp").Start()
		_ = cmd
	case "3":
		cmd := exec.Command("mpc","-q","toggle").Start()
		_ = cmd
	case "4":
		cmd := exec.Command("mpc","-q","next").Start()
		_ = cmd
	case "5":
		cmd := exec.Command("mpcprevious.sh").Start()
		_ = cmd
	}
}
