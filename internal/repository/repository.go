package repository

type AuthRepository interface {
}

type ChatRepository interface {
}

type Repository struct {
	AuthRepository
	ChatRepository
}

func New() *Repository {
	return &Repository{}
}
