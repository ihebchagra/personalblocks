package blocks

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"net/http"
	"io/ioutil"
	"strings"
)

func gettitle(radio string) string {
	var title string
	switch radio{
	case "normal":
		format := "[%artist% - ]|[%artist% - ]&%title%"
		song, _ := exec.Command("mpc","-f",format,"current").Output()
		title = string(song)
	case "pianorama":
		format := "%title%"
		song, _ := exec.Command("mpc","-f",format,"current").Output()
		title = string(song)
	case "misk":
		req, _ := http.NewRequest("GET", "http://www.misk.art/live", nil)
		res, _ := http.DefaultClient.Do(req)
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		lines := strings.Split(string(body),"\n")
		finishedartist := true
		finishedsong := true
		var artist,song string
		for i:=0; i <= len(lines) && (finishedartist || finishedsong); i++ {
			artistRe := regexp.MustCompile(`.*<.*pl-item-artist">(.+)<.*>.*`)
			songRe := regexp.MustCompile(`.*<.*pl-item-title">(.+)<.*>.*`)
			if artistRe.MatchString(lines[i]) == true && finishedartist {
				finishedartist = false
				artist = artistRe.ReplaceAllString(lines[i],"$1")
				artist = strings.Title(strings.ToLower(artist))
			}
			if songRe.MatchString(lines[i]) == true && finishedsong{
				finishedsong = false
				song = songRe.ReplaceAllString(lines[i],"$1")
				song = strings.Title(strings.ToLower(song))
			}
		}
		title = fmt.Sprintf("%s - %s\n",artist,song)
	}
	return title
}

func mpdradio(radio,button string) {
	switch button {
	case "":
		fmt.Printf("%s",gettitle(radio))
	case "2","4","5":
		title:=gettitle(radio)
		fmt.Printf("%s",title)
		home, _ := os.UserHomeDir()
		f, _ := os.OpenFile(home + "/.calcurse/notes/songstodownload.txt",os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		defer f.Close()
		if _, err := f.WriteString(title); err == nil {
			cmd := exec.Command("notify-send","🎶 Added to songstodownload.txt",title).Run()
			_ = cmd
		}
	}
}

func misk(radio,button string) {
	switch button {
	case "":
		fmt.Printf("%s",gettitle(radio))
	case "2","4","5":
		title:=gettitle(radio)
		fmt.Printf("%s",title)
		f, _ := os.OpenFile("text.log",os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		defer f.Close()
		if _, err := f.WriteString(title); err == nil {
			cmd := exec.Command("notify-send","🎶 Added to songstodownload.txt",title).Run()
			_ = cmd
		}
	}
}

func otherradiocommand(button,name,process string) {
	if button == "" {
		fmt.Printf("Playing %s\n",name)
	} else if button == "3" {
		cmd := exec.Command("pkill","-f",process)
		_ = cmd
	}
}


func other(radio,process,button string) {
	switch radio {
	case "exam":
		otherradiocommand(button,"Exam noise",process)
	case "thundernoise":
		otherradiocommand(button,"Thunder noise",process)
	case "brownnoise":
		otherradiocommand(button,"Brownian noise",process)
	case "lofimix":
		otherradiocommand(button,"LoFi HipHop mix",process)
	case "lofiradio":
		otherradiocommand(button,"LoFi HipHop radio",process)
	}
}


func Sounds(button string) {
	out, _ := exec.Command("mpc","current","-f","%file%").Output()
	file := string(out)
	lainRe := regexp.MustCompile(`^https:\/\/lainon\.life\/radio\/.*\.ogg\n$`)
	if file == "https://listen.moe/stream\n" || lainRe.MatchString(file) {
		mpdradio("normal",button)
	} else if file == "http://stream.pianoramaradio.ru\n" || lainRe.MatchString(file) {
		mpdradio("pianorama",button)
	} else if err := exec.Command("pgrep","-f","mpv http://live.misk.tn:8000/stream").Run(); err == nil {
		misk("misk",button)
	} else if err := exec.Command("pgrep","-f","examinationtime.mp3").Run(); err == nil {
		other("exam","examinationtime.mp3",button)
	} else if err := exec.Command("pgrep","-f","thundernoise.mp3").Run(); err == nil {
		other("thundernoise","thundernoise.mp3",button)
	} else if err := exec.Command("pgrep","-f","brownnoise.mp3").Run(); err == nil {
		other("brownnoise","brownnoise.mp3",button)
	} else if err := exec.Command("pgrep","-f","lofihiphopmix.mp3").Run(); err == nil {
		other("lofimix","lofihiphopmix.mp3",button)
	} else if err := exec.Command("pgrep","-f","https://www.youtube.com/watch\\?v=hHW1oY26kxQ").Run(); err == nil {
		other("lofiradio","https://www.youtube.com/watch?v=hHW1oY26kxQ",button)
	}
}
