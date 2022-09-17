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

func (q *Querier) GetUserById(ctx context.Context, id int) (entity.User, error) {
	return getUser(q, ctx, id)
}

func (q *Querier) UpdateUserPassword(ctx context.Context, userId int, password string) error {
	return updatePassword(q, ctx, userId, password)
}
