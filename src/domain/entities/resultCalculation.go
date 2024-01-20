package entities

type Action struct {
	ContentType string  `json:"contenttype"`
	ContentID   int     `json:"contentid"`
	Duration    float64 `json:"duration"`
}

type Data struct {
	ID     float64  `json:"_id"`
	Action []Action `json:"action"`
}
