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
	DeleteUser(ctx context.Context, obj *entity.User) error
}

type RDBSetRepo interface {
	RDBSet(ctx context.Context, RDBkey string, value interface{}, expiration time.Duration) error
}

type RDBGetRepo interface {
	RDBGet(ctx context.Context, RDBkey string) (string, error)
}

type FetchTagsRepo interface {
	FetchTags(ctx context.Context) ([]*entity.Tag, error)
}

type SaveTagRepo interface {
	SaveTag(ctx context.Context, obj *entity.Tag) error
}

type FindTagByTagRepo interface {
	FindTagByTag(ctx context.Context, tag string) (*entity.Tag, error)
}

type FindTagByIDRepo interface {
	FindTagByID(ctx context.Context, id int64) (*entity.Tag, error)
}

type DeleteTagRepo interface {
	DeleteTag(ctx context.Context, obj *entity.Tag) error
}

type SaveCategoryRepo interface {
	SaveCategory(ctx context.Context, obj *entity.Category) error
}

type FetchCategoriesRepo interface {
	FetchCategories(ctx context.Context) ([]*entity.Category, error)
}

type FindCategoryByCategoryRepo interface {
	FindCategoryByCategory(ctx context.Context, category string) (*entity.Category, error)
}

type FindCategoryByIDRepo interface {
	FindCategoryByID(ctx context.Context, id int64) (*entity.Category, error)
}

type DeleteCategoryRepo interface {
	DeleteCategory(ctx context.Context, obj *entity.Category) error
}

type FindPostBySlugRepo interface {
	FindPostBySlug(ctx context.Context, slug string) (*entity.Post, error)
}

type FindPostByIDRepo interface {
	FindPostByID(ctx context.Context, id int64) (*entity.Post, error)
}

type SavePostRepo interface {
	SavePost(ctx context.Context, obj *entity.Post) error
}

type CreatePostRepo interface {
	CreatePost(ctx context.Context, obj *entity.Post) error
}

type FindCategoriesByIDsRepo interface {
	FindCategoriesByIDs(ctx context.Context, ids []int64) ([]*entity.Category, error)
}

type FindTagsByIDsRepo interface {
	FindTagsByIDs(ctx context.Context, ids []int64) ([]*entity.Tag, error)
}

type FetchPostsRepo interface {
	FetchPosts(ctx context.Context) ([]*entity.Post, error)
}

type DeletePostRepo interface {
	DeletePost(ctx context.Context, obj *entity.Post) error
}

type FetchPostsByUserUsernameRepo interface {
	FetchPostsByUserUsername(ctx context.Context, username string) (*entity.User, error)
}
