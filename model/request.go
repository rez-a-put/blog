package model

type ReqDataPost struct {
	Id      string   `json:"-"`
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}

type ReqDataTag struct {
	Id    string `json:"-"`
	Label string `json:"label"`
}
