package blocks

import (
	"fmt"
	"os/exec"
	"strings"
	"strconv"
)


func Volume(button string) {
	out1, _ := exec.Command("pulsemixer","--get-mute").Output()
	out2, _ := exec.Command("pulsemixer","--get-volume").Output()
	if mute := string(out1); mute == "1\n" {
		fmt.Println("MUTE")
	} else {
		if vol, _ := strconv.Atoi(strings.Split(string(out2)," ")[0]); vol > 70 {
			fmt.Printf(" %d%%\n", vol)
		} else if vol > 30 {
			fmt.Printf(" %d%%\n", vol)
		} else {
			fmt.Printf(" %d%%\n", vol)
		}
	}


	switch button {
	case "1":
		cmd := exec.Command("i3-msg","-q","exec pavucontrol").Start()
		_ = cmd
	case "3":
		cmd := exec.Command("pulsemixer","--toggle-mute").Start()
		_ = cmd
	case "4":
		cmd := exec.Command("pulsemixer","--change-volume","+5").Start()
		_ = cmd
	case "5":
		cmd := exec.Command("pulsemixer","--change-volume","-5").Start()
		_ = cmd
	}
}
