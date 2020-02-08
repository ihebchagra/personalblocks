package blocks

import (
	"fmt"
	"os/exec"
	"strings"
)


func Mail(button string) {
	out, _:= exec.Command("sh","-c","du -a /home/iheb/.local/share/mail/*/INBOX/new/*").Output()
	if mail := strings.TrimSuffix(string(out),"\n"); mail != "" {
		fmt.Printf(" %d\n",len(strings.Split(mail,"\n")))
	}

	switch button {
	case "1":
		cmd := exec.Command("i3-msg","-q","exec $TERMINAL -e neomutt").Start()
		_ = cmd
	case "3":
		cmd := exec.Command("notify-send"," Syncing your mail").Start()
		cmd = exec.Command("mailsync").Start()
		_ = cmd
	}
}
