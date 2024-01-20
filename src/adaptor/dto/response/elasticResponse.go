package response

type ElasticResponse struct {
	UserID      interface{}
	ContentList []ContentDetail
}

type ContentDetail struct {
	Service     string `json:"service"`
	ContentType string `json:"content_type"`
	ContentId   int    `json:"content_id"`
}
