package postgres

import (
	"github.com/gofrs/uuid"
	"github.com/testing/go-kit-tutorial/model"
)

type DBReaderWriter interface {
	GetUser(id uuid.UUID) (interface{}, error)
	GetListUser(req model.ParamList) (model.Users, error)
	CreateUser(request interface{}) error
}
