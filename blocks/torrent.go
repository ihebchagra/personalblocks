package blocks

import (
	"fmt"
	"regexp"
	"os/exec"
	"strings"
)


func Torrent(button string) {
	out, _ := exec.Command("deluge-console","info").Output()
	info := string(out)
	lines := strings.Split(info,"\n")

	rePaused := regexp.MustCompile(`^State.*Paused.*`)
	reSeeding := regexp.MustCompile(`^State.*Seeding.*`)
	reQueued := regexp.MustCompile(`^State.*Queued.*`)
	reStalled := regexp.MustCompile(`^State.*Downloading Down Speed: 0\.0.*`)
	reDownloading := regexp.MustCompile(`^State.*Downloading.*`)

	var paused,seeding,queued,stalled,downloading int
	for i:=0; i < len(lines); i++ {
		if rePaused.MatchString(lines[i]) {
			paused++
		} else if reSeeding.MatchString(lines[i]) {
			seeding++
		} else if reQueued.MatchString(lines[i]) {
			queued++
		} else if reStalled.MatchString(lines[i]) {
			stalled++
		} else if reDownloading.MatchString(lines[i]) {
			downloading++
		}
	}

	var status string
	if paused > 0 {
		status += fmt.Sprintf("  %d",paused)
	}
	if seeding > 0 {
		status += fmt.Sprintf("  %d",seeding)
	}
	if queued > 0 {
		status += fmt.Sprintf("  %d",queued)
	}
	if stalled > 0 {
		status += fmt.Sprintf("  %d",stalled)
	}
	if downloading > 0 {
		status += fmt.Sprintf("  %d",downloading)
	}
	if status != "" {
		status += "\n"
		status = strings.TrimPrefix(status, " ")
	}

	fmt.Printf(status)



	if button == "1" {
		cmd := exec.Command("i3-msg","-q","exec deluge-gtk").Start()
		_ = cmd
	}
}
