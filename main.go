package main

import (
	"flag"
	"fmt"
	"time"
)

var (
	rawPlaylist = flag.String("playlist", "", "The youtube playlist id")
	rawDuration = flag.String("duration", (30 * time.Second).String(), "The duration of the clip")

	helpMsg = `
-playlist: the id of the youtube playlist.
-duration: the duration of the clip to select.`
)

func main() {
	flag.Parse()

	duration, err := time.ParseDuration(*rawDuration)
	if err != nil || *rawPlaylist == "" {
		fmt.Println(helpMsg)
		return
	}

	fmt.Printf("playlist: '%v'\n", *rawPlaylist)
	fmt.Printf("duration: '%v'\n", duration)
}
