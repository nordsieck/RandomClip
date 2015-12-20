package api

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

type Playlist struct {
	Kind     string   `json:"kind"`
	ETag     string   `json:"etag"`
	PageInfo PageInfo `json:"pageInfo"`
	Items    []Item   `json:"items"`
}

func (p *Playlist) IsZero() bool {
	return p.Kind == "" && p.ETag == "" && p.PageInfo.IsZero() && len(p.Items) == 0
}

func (p *Playlist) VideoIDs() []string {
	list := make([]string, 0, len(p.Items))
	for _, i := range p.Items {
		list = append(list, i.ContentDetails.VideoID)
	}
	return list
}

type PageInfo struct {
	TotalResults   uint8
	ResultsPerPage uint8
}

func (p *PageInfo) IsZero() bool { return p.TotalResults == 0 && p.ResultsPerPage == 0 }

type Item struct {
	Kind           string  `json:"kind"`
	ETag           string  `json:"etag"`
	ID             string  `json:"id"`
	ContentDetails Details `json:"contentDetails"`
}

type Details struct {
	VideoID string `json:"videoId"`
}

func DeserializePlaylist(body io.Reader) (*Playlist, error) {
	bodyBytes, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}

	var pl Playlist
	err = json.Unmarshal(bodyBytes, &pl)
	if err != nil {
		return nil, err
	}
	if pl.IsZero() {
		var e Error
		err = json.Unmarshal(bodyBytes, &e)
		if err != nil {
			return nil, err
		}
		if e.IsZero() {
			return nil, ErrDeserializeInput
		}
		return nil, &e
	}
	return &pl, nil
}
