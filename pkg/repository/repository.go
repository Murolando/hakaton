package repository

import (
	"github.com/Murolando/hakaton_geo/ent"
	"github.com/Murolando/hakaton_geo/pkg/repository/postgres"
	"github.com/jmoiron/sqlx"
)

type Auth interface {
	SignUp(user ent.User) (int64, error)
	GetUserByLoginAndPassword(mail *string, password *string) (int64, error)
	SetSession(user int64, refresh string, expiredAt string) error
	GetByRefreshToken(refresh string) (int64, error)
}
type Class interface{
	DashboardClass(userId int64) ([]*ent.ClassProgressResponce, error)
	CommonProgressInfo(userId int64)(*ent.ChildDashClassResponce,error)

	MyClass(userId int64)([]*ent.ChildMyClassResponce,error)
}
type Repository struct {
	Class
	Auth
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Auth: postgres.NewAuthPostgres(db),
		Class: postgres.NewClassPostgres(db),
	}
}
