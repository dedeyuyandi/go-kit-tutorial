package service

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gofrs/uuid"
)

func (s service) GetUser(ctx context.Context, id uuid.UUID) (interface{}, error) {
	logger := log.With(s.logger, "method", "GetUser")

	resp, err := s.repostory.GetUser(id)

	if err != nil {
		level.Error(logger).Log("err", err)
		return nil, err
	}

	logger.Log("Get user", id)

	return resp, nil
}
