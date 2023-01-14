package models

type Dates struct {
	Start string `json:"date_start,omitempty" `
	End   string `json:"date_end,omitempty" `
}

type Tweet struct {
	Data []*InfoTweet `json:"data,omitempty" `
	Meta *Meta        `json:"meta,omitempty" `
}

type Meta struct {
	NextToken   string `json:"next_token, omitempty"`
	ResultCount int    `json:"result_count, omitempty"`
}

type InfoTweet struct {
	Text       string `json:"text, omitempty"`
	Id         string `json:"id, omitempty"`
	Created_at string `json:"created_at, omitempty"`
}
