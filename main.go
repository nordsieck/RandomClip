package main

import (
	"flag"
	"fmt"
)

var (
	rawPlaylist = flag.String("playlist", "", "The youtube playlist id")
	rawDuration = flag.String("duration", "30s", "The duration of the clip")
)

func main() {
	flag.Parse()

	fmt.Printf("playlist: '%v'\n", *rawPlaylist)
	fmt.Printf("duration: '%v'\n", *rawDuration)
}
