package entities

type ListBehaviour struct {
	UserID  interface{} `json:"user_id"`
	Content []Content   `json:"content"`
}

type Content struct {
	ContentID  int    `json:"content_id"`
	ContetType string `json:"contet_type"`
}
