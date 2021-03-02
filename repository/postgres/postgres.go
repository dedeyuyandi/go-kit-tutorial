package postgres

import (
	"database/sql"

	"github.com/go-kit/kit/log"
)

var (
	createUser  = `INSERT INTO users (id, email, password) VALUES ($1, $2, $3)`
	getUserByID = `SELECT email, password FROM users WHERE id=$1`
	getListUser = `SELECT id, email, password FROM users ORDER BY $1 LIMIT $2,$3`
)

type repo struct {
	db     *sql.DB
	logger log.Logger
}

func NewRepo(db *sql.DB, logger log.Logger) DBReaderWriter {
	return &repo{
		db:     db,
		logger: log.With(logger, "repo", "sql"),
	}
}
