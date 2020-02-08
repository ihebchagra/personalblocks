package main

import (
	"fmt"
	"os"
	"github.com/ihebchagra/personalblocks/blocks"
)

func main() {
	help := "Choose a command from:\n" +
	"-battery\n" +
	"-mail\n" +
	"-music\n" +
	"-news\n" +
	"-nmblocks\n" +
	"-root\n" +
	"-screenawake\n" +
	"-screencast\n" +
	"-sounds\n" +
	"-time\n" +
	"-torrent\n" +
	"-volume"
	button := os.Getenv("BLOCK_BUTTON")
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "battery":
			blocks.Battery(button)
		case "mail":
			blocks.Mail(button)
		case "music":
			blocks.Music(button)
		case "news":
			blocks.News(button)
		case "nmblocks":
			blocks.Nmblocks(button)
		case "root":
			blocks.Root(button)
		case "screenawake":
			blocks.Screenawake(button)
		case "screencast":
			blocks.Screencast(button)
		case "sounds":
			blocks.Sounds(button)
		case "time":
			blocks.Time(button)
		case "torrent":
			blocks.Torrent(button)
		case "volume":
			blocks.Volume(button)
		default:
			fmt.Println(help)
		}
	} else {
		fmt.Println(help)
	}
}
