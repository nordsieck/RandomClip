package api

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/nordsieck/defect"
)

const playlistExample = `{
 "kind": "youtube#playlistItemListResponse",
 "etag": "\"kuL0kDMAqRo3pU7O0pwlO-Lfzp4/i1ORKHHBLiGYbgwmVrbz1HTnjiM\"",
 "pageInfo": {
  "totalResults": 4,
  "resultsPerPage": 50
 },
 "items": [
  {
   "kind": "youtube#playlistItem",
   "etag": "\"kuL0kDMAqRo3pU7O0pwlO-Lfzp4/pxjefDqZZB28ihx1tU6Lacs3OWI\"",
   "id": "PLTNzmeZSD6S8fiu_sThS48yscKcogb0k1XR5NFcpZmTQ",
   "contentDetails": {
    "videoId": "WSrktmE963I"
   }
  },
  {
   "kind": "youtube#playlistItem",
   "etag": "\"kuL0kDMAqRo3pU7O0pwlO-Lfzp4/xz1RoE48BdkfF8HOBj9a1KxpYMc\"",
   "id": "PLTNzmeZSD6S8fiu_sThS48w1IS2M_ryLbIL_lPto-mTk",
   "contentDetails": {
    "videoId": "9Sc-ir2UwGU"
   }
  },
  {
   "kind": "youtube#playlistItem",
   "etag": "\"kuL0kDMAqRo3pU7O0pwlO-Lfzp4/yXbzH0XnN3u8VaUs7mlOIYVm8DE\"",
   "id": "PLTNzmeZSD6S8fiu_sThS48_DAMTUiMurXx53dT5nYfRM",
   "contentDetails": {
    "videoId": "3gPknJGnmUw"
   }
  },
  {
   "kind": "youtube#playlistItem",
   "etag": "\"kuL0kDMAqRo3pU7O0pwlO-Lfzp4/r1OcCvTmqZsPaCN87eRmaCmrUQY\"",
   "id": "PLTNzmeZSD6S8fiu_sThS48xP4kf4RUpRQKSduLV2-2SY",
   "contentDetails": {
    "videoId": "PCcxMQoHznA"
   }
  }
 ]
}`

var expectedPlaylist = Playlist{
	Kind:     "youtube#playlistItemListResponse",
	ETag:     "\"kuL0kDMAqRo3pU7O0pwlO-Lfzp4/i1ORKHHBLiGYbgwmVrbz1HTnjiM\"",
	PageInfo: PageInfo{TotalResults: 4, ResultsPerPage: 50},
	Items: []Item{{
		Kind:           "youtube#playlistItem",
		ETag:           "\"kuL0kDMAqRo3pU7O0pwlO-Lfzp4/pxjefDqZZB28ihx1tU6Lacs3OWI\"",
		ID:             "PLTNzmeZSD6S8fiu_sThS48yscKcogb0k1XR5NFcpZmTQ",
		ContentDetails: Details{VideoID: "WSrktmE963I"},
	}, {
		Kind:           "youtube#playlistItem",
		ETag:           "\"kuL0kDMAqRo3pU7O0pwlO-Lfzp4/xz1RoE48BdkfF8HOBj9a1KxpYMc\"",
		ID:             "PLTNzmeZSD6S8fiu_sThS48w1IS2M_ryLbIL_lPto-mTk",
		ContentDetails: Details{VideoID: "9Sc-ir2UwGU"},
	}, {
		Kind:           "youtube#playlistItem",
		ETag:           "\"kuL0kDMAqRo3pU7O0pwlO-Lfzp4/yXbzH0XnN3u8VaUs7mlOIYVm8DE\"",
		ID:             "PLTNzmeZSD6S8fiu_sThS48_DAMTUiMurXx53dT5nYfRM",
		ContentDetails: Details{VideoID: "3gPknJGnmUw"},
	}, {
		Kind:           "youtube#playlistItem",
		ETag:           "\"kuL0kDMAqRo3pU7O0pwlO-Lfzp4/r1OcCvTmqZsPaCN87eRmaCmrUQY\"",
		ID:             "PLTNzmeZSD6S8fiu_sThS48xP4kf4RUpRQKSduLV2-2SY",
		ContentDetails: Details{VideoID: "PCcxMQoHznA"},
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
