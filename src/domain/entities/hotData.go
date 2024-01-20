package entities

type HotContentTypeVideo struct {
	ContentId   int    `json:"content_id"`
	Title       string `json:"title"`
	Thumbnail   string `json:"thumbnail"`
	UrlVideo    string `json:"url_video"`
	Description string `json:"description"`
	User        struct {
		UserId         int    `json:"user_id"`
		DisplayName    string `json:"display_name"`
		TotalFollowers int    `json:"total_followers"`
		IsFollow       bool   `json:"is_follow"`
		Avatar         string `json:"avatar"`
	} `json:"user"`
	Competition struct {
		Id             int         `json:"id"`
		Name           string      `json:"name"`
		ProgramType    string      `json:"program_type"`
		Genre          interface{} `json:"genre"`
		Classification string      `json:"classification"`
	} `json:"competition"`
	IsVote       bool          `json:"is_vote"`
	VoteTimer    int           `json:"vote_timer"`
	LikeCount    int           `json:"like_count"`
	IsLike       bool          `json:"is_like"`
	CommentCount int           `json:"comment_count"`
	UrlShare     string        `json:"url_share"`
	Hashtags     []interface{} `json:"hashtags"`
}
