package service

import (
	"github.com/Murolando/hakaton_geo/ent"
	"github.com/Murolando/hakaton_geo/pkg/repository"
	repositoryImage "github.com/Murolando/hakaton_geo/pkg/repository_image"
)

type Auth interface {
	SignIn(mail *string, password *string) (int64, error)
	SignUp(user ent.User) (map[string]interface{}, error)
	GenerateToken(id int64) (string, error)
	ParseToken(accesstoken string) (int64, error)
	NewRefreshToken(id int64) (string, error)
	GetByRefreshToken(refresh string) (int64, error)
}

type Service struct {
	Auth
}

func NewService(repo *repository.Repository, image *repositoryImage.Image) *Service {
	return &Service{
		Auth: NewAuthService(repo),
	}
}
