package entities

type NewsContentTypeHtml struct {
	ContentId   int           `json:"content_id"`
	ContentType string        `json:"content_type"`
	Title       string        `json:"title"`
	Category    string        `json:"category"`
	PubDate     int           `json:"pub_date"`
	ImageUrl    string        `json:"image_url"`
	Content     string        `json:"content"`
	Source      string        `json:"source"`
	Permalink   string        `json:"permalink"`
	LikeCount   int           `json:"like_count"`
	IsLike      bool          `json:"is_like"`
	Tags        []interface{} `json:"tags"`
}
