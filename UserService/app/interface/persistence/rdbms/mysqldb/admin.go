package mysqldb

import (
	entity "SepFirst/UserService/app/domain/entities"
	"context"
)

func (q *Querier) GetAllUsers(ctx context.Context) ([]entity.User, error) {
	return getAllUsers(q, ctx)
}

func (q *Querier) DeleteUser(ctx context.Context, userId int) error {
	return deleteUser(q, ctx, userId)
}

func (q *Querier) BlockUser(ctx context.Context, userId int) error {
	return blockUser(q, ctx, userId)
}
