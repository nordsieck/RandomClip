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

	MaxPlaylistItems = "50" // max value allowed: https://developers.google.com/youtube/v3/docs/playlistItems/list
)

func GetPlaylistFragment(key, playlist, token string) (videoIDs []string, pageToken string, e error) {
	params := url.Values{
		Key:        {key},
		Part:       {ContentDetails},
		PlaylistID: {playlist},
		MaxResults: {MaxPlaylistItems},
	}
	if token != "" {
		params[PageToken] = []string{token}
	}
	resp, err := http.Get(PlaylistItems + "?" + params.Encode())
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()

	pl, err := DeserializePlaylist(resp.Body)
	if err != nil {
		return nil, "", err
	}

	return pl.VideoIDs(), pl.NextPageToken, nil
}

func PlaylistVideos(key, playlist string) ([]string, error) {
	fullList := []string{}

	token := ""
	for {
		var list []string
		var err error
		list, token, err = GetPlaylistFragment(key, playlist, token)
		if err != nil {
			return nil, err
		}
		fullList = append(fullList, list...)
		if token == "" {
			break
		}
	}

	return fullList, nil
}
