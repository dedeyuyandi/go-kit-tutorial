package postgres

import (
	"errors"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/testing/go-kit-tutorial/model"
)

func (repo *repo) GetUser(ID uuid.UUID) (interface{}, error) {
	row := repo.db.QueryRow(getUserByID, ID)
	var u model.GetUserResponse
	if err := row.Scan(
		&u.Email,
		&u.Password,
	); err != nil {
		fmt.Println(err)
		return nil, errors.New("ErrDataNotFound")
	}

	return u, nil
}
