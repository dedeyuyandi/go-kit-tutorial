package model

import "github.com/gofrs/uuid"

type ParamList struct {
	Offset  int32
	Limit   int32
	OrderBy string
}

func (tc ParamList) IsEmpty() bool {
	if tc.Offset < 0 || tc.Limit == 0 || tc.OrderBy == "" {
		return false
	}
	return true
}

type (
	CreateUserResponse struct {
		Message string `json:"message"`
	}

	GetUserRequest struct {
		Id uuid.UUID `json:"id"`
	}

	GetUserResponse struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)

type WriteUserRequest struct {
	ID       uuid.UUID `json:"id,omitempty"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

func (tc WriteUserRequest) IsEmpty() bool {
	if tc.Email == "" || tc.Password == "" {
		return false
	}
	return true
}

type User struct {
	WriteUserRequest
}

type Users []WriteUserRequest
