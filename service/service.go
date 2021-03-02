package service

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/gofrs/uuid"
	"github.com/testing/go-kit-tutorial/model"
	"github.com/testing/go-kit-tutorial/repository/postgres"
)

type Service interface {
	CreateUser(ctx context.Context, req model.WriteUserRequest) error
	GetUser(ctx context.Context, id uuid.UUID) (interface{}, error)
}

type service struct {
	repostory postgres.DBReaderWriter
	logger    log.Logger
}

func NewService(rep postgres.DBReaderWriter, logger log.Logger) Service {
	return &service{
		repostory: rep,
		logger:    logger,
	}
}
