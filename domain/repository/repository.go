package repository

import (
	"context"
	"time"

	"github.com/raismaulana/blogP/domain/entity"
)

type SaveUserRepo interface {
	SaveUser(ctx context.Context, obj *entity.User) error
}

type FindUserByUsernameRepo interface {
	FindUserByUsername(ctx context.Context, username string, scope bool) (*entity.User, error)
}

type FindUserByEmailRepo interface {
	FindUserByEmail(ctx context.Context, email string, scope bool) (*entity.User, error)
}

type FindUserByIDRepo interface {
	FindUserByID(ctx context.Context, ID int64, scope bool) (*entity.User, error)
}

type FetchUsersRepo interface {
	FetchUsers(ctx context.Context, scope bool) ([]*entity.User, error)
}

type DeleteUserRepo interface {
	DeleteUser(ctx context.Context, ID *entity.User) error
}

type RDBSetRepo interface {
	RDBSet(ctx context.Context, RDBkey string, value interface{}, expiration time.Duration) error
}

type RDBGetRepo interface {
	RDBGet(ctx context.Context, RDBkey string) (string, error)
}

type FetchTagsRepo interface {
	FetchTags(ctx context.Context, scope bool) ([]*entity.Tag, error)
}
