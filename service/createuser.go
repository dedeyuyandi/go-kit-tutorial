package service

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/testing/go-kit-tutorial/model"
)

func (s service) CreateUser(ctx context.Context, req model.WriteUserRequest) error {
	logger := log.With(s.logger, "method", "CreateUser")

	if err := s.repostory.CreateUser(req); err != nil {
		level.Error(logger).Log("err", err)
		return err
	}

	return nil
}
