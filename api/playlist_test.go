package api

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/nordsieck/defect"
)

const playlistExample = `{
 "kind": "youtube#playlistItemListResponse",
 "etag": "\"kuL0kDMAqRo3pU7O0pwlO-Lfzp4/npdqJe-4vvFRFapzc8fzhv0TW-8\"",
 "nextPageToken": "CAgQAA",
 "prevPageToken": "CAQQAQ",
 "pageInfo": {
  "totalResults": 78,
  "resultsPerPage": 4
 },
 "items": [
  {
   "kind": "youtube#playlistItem",
   "etag": "\"kuL0kDMAqRo3pU7O0pwlO-Lfzp4/ya8nbnto-i637Rin-EyoqbKg-vI\"",
   "id": "PLox-BzA6rjaZF4rFbXwqaC7vpIq5hmYL0_017yrOjtrw",
   "contentDetails": {
    "videoId": "h_hOQAZHy1A"
   }
  },
  {
   "kind": "youtube#playlistItem",
   "etag": "\"kuL0kDMAqRo3pU7O0pwlO-Lfzp4/EFau6eavCwIvHjYMq4rYSJjbkAA\"",
   "id": "PLox-BzA6rjaZF4rFbXwqaC4tnYN92LGpTcN3LZFkgS0g",
   "contentDetails": {
    "videoId": "FUak2C_KEeU"
   }
  },
  {
   "kind": "youtube#playlistItem",
   "etag": "\"kuL0kDMAqRo3pU7O0pwlO-Lfzp4/twrX_a-7puSBmUACFUmGBybhbCQ\"",
   "id": "PLox-BzA6rjaZF4rFbXwqaC3bnDWpreH-2b2H8pfQ9Pgw",
   "contentDetails": {
    "videoId": "D6DFLNa6MBA"
   }
  },
  {
   "kind": "youtube#playlistItem",
   "etag": "\"kuL0kDMAqRo3pU7O0pwlO-Lfzp4/_2Cral9lBX6QegayYW7B1gTVxFs\"",
   "id": "PLox-BzA6rjaZF4rFbXwqaC8NShPDo3RtU_xbBUgMJv3A",
   "contentDetails": {
    "videoId": "gCJ3rmiZFr8"
   }
  }
 ]
}`

var expectedPlaylist = Playlist{
	Kind:          "youtube#playlistItemListResponse",
	ETag:          "\"kuL0kDMAqRo3pU7O0pwlO-Lfzp4/npdqJe-4vvFRFapzc8fzhv0TW-8\"",
	NextPageToken: "CAgQAA",
	PrevPageToken: "CAQQAQ",
	PageInfo:      PageInfo{TotalResults: 78, ResultsPerPage: 4},
	Items: []Item{{
		Kind:           "youtube#playlistItem",
		ETag:           "\"kuL0kDMAqRo3pU7O0pwlO-Lfzp4/ya8nbnto-i637Rin-EyoqbKg-vI\"",
		ID:             "PLox-BzA6rjaZF4rFbXwqaC7vpIq5hmYL0_017yrOjtrw",
		ContentDetails: Details{VideoID: "h_hOQAZHy1A"},
	}, {
		Kind:           "youtube#playlistItem",
		ETag:           "\"kuL0kDMAqRo3pU7O0pwlO-Lfzp4/EFau6eavCwIvHjYMq4rYSJjbkAA\"",
		ID:             "PLox-BzA6rjaZF4rFbXwqaC4tnYN92LGpTcN3LZFkgS0g",
		ContentDetails: Details{VideoID: "FUak2C_KEeU"},
	}, {
		Kind:           "youtube#playlistItem",
		ETag:           "\"kuL0kDMAqRo3pU7O0pwlO-Lfzp4/twrX_a-7puSBmUACFUmGBybhbCQ\"",
		ID:             "PLox-BzA6rjaZF4rFbXwqaC3bnDWpreH-2b2H8pfQ9Pgw",
		ContentDetails: Details{VideoID: "D6DFLNa6MBA"},
	}, {
		Kind:           "youtube#playlistItem",
		ETag:           "\"kuL0kDMAqRo3pU7O0pwlO-Lfzp4/_2Cral9lBX6QegayYW7B1gTVxFs\"",
		ID:             "PLox-BzA6rjaZF4rFbXwqaC8NShPDo3RtU_xbBUgMJv3A",
		ContentDetails: Details{VideoID: "gCJ3rmiZFr8"},
	}},
}

func TestPlaylistUnmarshal(t *testing.T) {
	var p Playlist
	err := json.Unmarshal([]byte(playlistExample), &p)
	defect.Equal(t, err, nil)
	defect.DeepEqual(t, p, expectedPlaylist)
}

func TestPlaylist_IsZero(t *testing.T) {
	p := Playlist{}
	defect.Equal(t, p.IsZero(), true)
	p.Kind = "kind"
	defect.Equal(t, p.IsZero(), false)
}

func TestDeserializePlaylist(t *testing.T) {
	buff := bytes.NewBuffer([]byte(playlistExample))

	pl, err := DeserializePlaylist(buff)
	defect.Equal(t, err, nil)
	defect.DeepEqual(t, pl, &expectedPlaylist)
}
