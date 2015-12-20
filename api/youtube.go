package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	PlaylistItems = "https://www.googleapis.com/youtube/v3/playlistItems"

	Key            = "key"
	Part           = "part"
	ContentDetails = "contentDetails"
	PlaylistID     = "playlistId"
	MaxResults     = "maxResults"
	PageToken      = "pageToken"
)

func PlaylistVideos(key, playlist string) ([]string, error) {
	params := url.Values{
		Key:        {key},
		Part:       {ContentDetails},
		PlaylistID: {playlist},
		MaxResults: {"50"}, // max value allowed: https://developers.google.com/youtube/v3/docs/playlistItems/list
	}
	resp, err := http.Get(PlaylistItems + "?" + params.Encode())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(body))

	return nil, nil
}
