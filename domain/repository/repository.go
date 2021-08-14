package repository

import (
	"context"

	"github.com/raismaulana/blogP/domain/entity"
)

type SaveUserRepo interface {
	SaveUser(ctx context.Context, obj *entity.User) error
}

type FindUserByUsernameRepo interface {
	FindUserByUsername(ctx context.Context, username string) (*entity.User, error)
}

type FindUserByEmailRepo interface {
	FindUserByEmail(ctx context.Context, email string) (*entity.User, error)
}

type FindUserByIDRepo interface {
	FindUserByID(ctx context.Context, ID int64) (*entity.User, error)
}

type FetchUsersRepo interface {
	FetchUsers(ctx context.Context) ([]*entity.User, error)
}

type DeleteUserRepo interface {
	DeleteUser(ctx context.Context, ID *entity.User) error
}
