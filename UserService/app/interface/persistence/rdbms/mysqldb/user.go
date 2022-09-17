package mysqldb

import (
	entity "SepFirst/UserService/app/domain/entities"
	"context"
)

func (q *Querier) CreateUser(ctx context.Context, user entity.User, password string) (int, error) {
	err := insertPassword(q, ctx, user.ID, password)
	if err != nil {
		return 0, err
	}

	return insertUser(q, ctx, user)
}
