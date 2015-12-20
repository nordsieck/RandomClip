package api

import (
	"net/http"
	"net/url"

	"github.com/nordsieck/defect"
)

const (
	PlaylistItems = "https://www.googleapis.com/youtube/v3/playlistItems"

	Key            = "key"
	Part           = "part"
	ContentDetails = "contentDetails"
	PlaylistID     = "playlistId"
	MaxResults     = "maxResults"
	PageToken      = "pageToken"

	ErrDeserializeInput = defect.Error("Unable to deserialize input")
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

	pl, err := DeserializePlaylist(resp.Body)
	if err != nil {
		return nil, err
	}
	return pl.VideoIDs(), nil
}
