package api

type Playlist struct {
	Kind     string   `json:"kind"`
	ETag     string   `json:"etag"`
	PageInfo PageInfo `json:"pageInfo"`
	Items    []Item   `json:"items"`
}

func (p *Playlist) IsZero() bool {
	return p.Kind == "" && p.ETag == "" && p.PageInfo.IsZero() && len(p.Items) == 0
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
