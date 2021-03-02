package postgres

import (
	"github.com/testing/go-kit-tutorial/model"
)

func (repo *repo) GetListUser(req model.ParamList) (model.Users, error) {
	rows, err := repo.db.Query(getListUser, req.OrderBy, req.Limit, req.Offset)
	if err != nil {
		return model.Users{}, err
	}
	defer rows.Close()

	result := make(model.Users, 0)
	for rows.Next() {
		t := model.WriteUserRequest{}
		err = rows.Scan(
			&t.ID,
			&t.Email,
			&t.Password,
		)
		if err != nil {
			return model.Users{}, err
		}

		result = append(result, t)
	}

	return result, nil
}
