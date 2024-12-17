package storage

// Storage интерфейс для работы с данными
type Storage interface {
	GetMessage() string
}

type storage struct{}

func NewStorage() Storage {
	return &storage{}
}

func (s *storage) GetMessage() string {
	return "Hello from the Go web server!"
}
