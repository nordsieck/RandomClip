package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/nordsieck/randomclip/api"
)

var (
	rawPlaylist = flag.String("playlist", "", "The youtube playlist id")
	rawDuration = flag.String("duration", (30 * time.Second).String(), "The duration of the clip")
	rawKey      = flag.String("key", "", "Your Google Developers API key")

	helpMsg = `
-playlist: the id of the youtube playlist.
-duration: the duration of the clip to select.`
)

func main() {
	flag.Parse()

	duration, err := time.ParseDuration(*rawDuration)
	if err != nil || *rawPlaylist == "" || *rawPlaylist == "" || *rawKey == "" {
		fmt.Println(helpMsg)
		return
	}

	_ = duration

	fmt.Println(api.PlaylistVideos(*rawKey, *rawPlaylist))
}
