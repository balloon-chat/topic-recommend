package model

type Topic struct {
	Id        string `json:"id"`
	CreatedAt int    `json:"createdAt"`
}

type TopicData struct {
	Topic
	MessageCount int
}
