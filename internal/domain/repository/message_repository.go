package repository

type MessageDatabase interface {
	GetMessageCountOf(topicId string) (*int, error)
}
