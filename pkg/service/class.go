package service

import (
	"github.com/Murolando/hakaton_geo/ent"
	"github.com/Murolando/hakaton_geo/pkg/repository"
	repositoryImage "github.com/Murolando/hakaton_geo/pkg/repository_image"
)

type ClassService struct {
	repo      *repository.Repository
	repoImage *repositoryImage.Image
}

func NewClassService(repo *repository.Repository, repoImage *repositoryImage.Image) *ClassService {
	return &ClassService{
		repo:      repo,
		repoImage: repoImage,
	}
}

func (s *ClassService) AllClass() ([]*ent.ClassResponce, error) {
	class,err := s.repo.AllClass()
	if err!=nil{
		return nil,err
	}
	return class,nil
}
