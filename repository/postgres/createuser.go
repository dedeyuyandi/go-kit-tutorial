package postgres

import (
	"errors"

	"github.com/gofrs/uuid"
	"github.com/micro/go-micro/logger"
	"github.com/testing/go-kit-tutorial/model"
)

func (repo *repo) CreateUser(request interface{}) error {
	req := request.(model.WriteUserRequest)

	if !req.IsEmpty() {
		return errors.New("Unable to handle Repo Request")
	}

	stmt, err := repo.db.Prepare(createUser)
	if err != nil {
		logger.Error(err)
		return errors.New("ErrStatement")
	}
	defer stmt.Close()

	id, _ := uuid.NewV4()
	req.ID = id

	_, err = stmt.Exec(req.ID, req.Email, req.Password)
	if err != nil {
		return errors.New("ErrExecuteDatabaseStatement")
	}
	return nil
}
