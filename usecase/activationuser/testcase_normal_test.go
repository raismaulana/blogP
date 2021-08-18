package activationuser

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/raismaulana/blogP/application/apperror"
	"github.com/raismaulana/blogP/domain/entity"
	"github.com/raismaulana/blogP/infrastructure/log"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

var (
	contextDBValue = "DB"
	db             = *&gorm.DB{
		Config:       &gorm.Config{},
		Error:        nil,
		RowsAffected: 0,
		Statement:    &gorm.Statement{},
	}
	users = []entity.User{
		{
			ID:       1,
			Username: "user1",
			Name:     "user1",
			Email:    "user1@gmail.com",
			Password: "user1",
			City:     "jakarta",
			Country:  "indonesia",
			Birthday: time.Date(1999, time.June, 25, 0, 0, 0, 0, time.UTC),
			WebProfile: null.String{
				NullString: sql.NullString{
					String: "user1.com",
					Valid:  true,
				},
			},
			ActivatedAt: null.Time{
				NullTime: sql.NullTime{
					Time:  time.Time{},
					Valid: false,
				},
			},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: gorm.DeletedAt{
				Time:  time.Time{},
				Valid: false,
			},
		},
		{
			ID:       2,
			Username: "user2",
			Name:     "user2",
			Email:    "user2@gmail.com",
			Password: "user2",
			City:     "jakarta",
			Country:  "indonesia",
			Birthday: time.Date(1999, time.June, 25, 0, 0, 0, 0, time.UTC),
			WebProfile: null.String{
				NullString: sql.NullString{
					String: "user2.com",
					Valid:  true,
				},
			},
			ActivatedAt: null.Time{
				NullTime: sql.NullTime{
					Time:  time.Now(),
					Valid: true,
				},
			},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: gorm.DeletedAt{
				Time:  time.Time{},
				Valid: false,
			},
		},
	}
)

type mockOutportNormal struct {
	t *testing.T
}

func (r *mockOutportNormal) FindUserByID(ctx context.Context, ID int64, scope bool) (*entity.User, error) {
	log.Info(ctx, "called")

	for _, v := range users {
		if v.ID == ID {
			return &v, nil
		}
	}

	return &entity.User{}, apperror.ObjectNotFound.Var(entity.User{})
}

func (r *mockOutportNormal) RDBGet(ctx context.Context, RDBkey string) (string, error) {
	log.Info(ctx, "called")

	return "thisisactivationcode", nil
}

func (r *mockOutportNormal) GetDatabase(ctx context.Context) (context.Context, error) {
	log.Info(ctx, "called")

	return context.WithValue(ctx, contextDBValue, db), nil
}

func (r *mockOutportNormal) SaveUser(ctx context.Context, obj *entity.User) error {
	log.Info(ctx, "called")

	return nil
}

func (r *mockOutportNormal) BeginTransaction(ctx context.Context) (context.Context, error) {
	log.Info(ctx, "called")

	return context.WithValue(ctx, contextDBValue, db.Begin()), nil
}

func (r *mockOutportNormal) CommitTransaction(ctx context.Context) error {
	log.Info(ctx, "called")

	return nil
}

func (r *mockOutportNormal) RollbackTransaction(ctx context.Context) error {
	log.Info(ctx, "called")

	return nil
}

// TestCaseNormal is for the case where the usecase activationuser works normally
// user1 tries to activate his account. He sends request (id, email, activation code) which is same with data in rdbms. His activation code is still in redis db.
func TestCaseNormal(t *testing.T) {

	ctx := context.Background()

	mockOutport := mockOutportNormal{
		t: t,
	}

	res, err := NewUsecase(&mockOutport).Execute(ctx, InportRequest{
		ID:             1,
		Email:          "user1@gmail.com",
		ActivationCode: "thisisactivationcode",
	})

	if err != nil {
		t.Errorf("%v", err.Error())
		t.FailNow()
	}

	t.Logf("%v", res)

}
