package postgres

import "github.com/jmoiron/sqlx"

const (
	usersTable      = "users"
	chatsTable      = "chats"
	usersChatsTable = "users_chats"
	messagesTable   = "message"
)

func New(dsn string) (*sqlx.DB, error) {

	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
