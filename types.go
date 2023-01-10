package main

type requestOptions struct {
	id           string
	maxImages    int
	file         string
	pretty       bool
	time         int
	base64images bool
	headers      map[string]string
}

type Item struct {
	ID       string  `json:"id"`
	Time     float64 `json:"time"`
	ImageURL string  `json:"imageUrl"`
	Likes    float64 `json:"likes"`
	Comments float64 `json:"comments" `
	Link     string  `json:"link" `
	Text     string  `json:"text"`
	Image    string  `json:"image,omitempty"`
}
