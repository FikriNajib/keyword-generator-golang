package entities

type VideoContentTypeClip struct {
	ContentId   int    `json:"content_id"`
	ContentType string `json:"content_type"`
	Description string `json:"description"`
	MediaUrl    string `json:"media_url"`
	Permalink   string `json:"permalink"`
	CreatedAt   int    `json:"created_at"`
	UpdatedAt   int    `json:"updated_at"`
	Program     struct {
		Id    int    `json:"id"`
		Title string `json:"title"`
		Image string `json:"image"`
	} `json:"program"`
	LikeCount int    `json:"like_count"`
	ShareUrl  int    `json:"share_url"`
	IsLike    bool   `json:"is_like"`
	MediaType string `json:"media_type"`
}

type VideoContentTypeStory struct {
	ContentId   int    `json:"content_id"`
	ContentType string `json:"content_type"`
	Title       string `json:"title"`
	Description string `json:"description"`
	MediaUrl    string `json:"media_url"`
	Permalink   string `json:"permalink"`
	CreatedAt   int    `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	Program     struct {
		Id             int    `json:"id"`
		Title          string `json:"title"`
		Image          string `json:"image"`
		Classification string `json:"classification"`
	} `json:"program"`
	Tv struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"tv"`
	LikeCount int    `json:"like_count"`
	ShareUrl  string `json:"share_url"`
	IsLike    bool   `json:"is_like"`
	MediaType string `json:"media_type"`
}
