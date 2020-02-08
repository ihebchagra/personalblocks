package blocks

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)


func Battery(button string) {
	out, _ := exec.Command("acpi","-b").CombinedOutput()
	acpi := strings.TrimSuffix(string(out),"\n")
	re1 := regexp.MustCompile(`^\w+ \w+: (\w+).*`)
	re2 := regexp.MustCompile(`.*, ([[:digit:]]+)%.*`)
	status := re1.ReplaceAllString(acpi ,"$1")
	capacity := re2.ReplaceAllString(acpi ,"$1")
	capacityval, _ := strconv.Atoi(capacity)
	if status == "Charging" {
		fmt.Printf(" %s%%\n",capacity)
	} else if capacityval > 75 {
		fmt.Printf(" %s%%\n",capacity)
	} else if capacityval > 50 {
		fmt.Printf(" %s%%\n",capacity)
	} else if capacityval > 20 {
		fmt.Printf(" %s%%\n",capacity)
	} else {
		fmt.Printf(" %s%% \n",capacity)
	}

	if button == "1" {
		cmd := exec.Command("notify-send",acpi).Start()
		_ = cmd
	}
}
