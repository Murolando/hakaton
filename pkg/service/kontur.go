package service

import (
	"github.com/Murolando/hakaton_geo/ent"
	"github.com/Murolando/hakaton_geo/pkg/repository"
	repositoryImage "github.com/Murolando/hakaton_geo/pkg/repository_image"
)

type KonturService struct {
	repo      *repository.Repository
	repoImage *repositoryImage.Image
}

func NewKonturService(repo *repository.Repository, repoImage *repositoryImage.Image) *KonturService {
	return &KonturService{
		repo:      repo,
		repoImage: repoImage,
	}
}

func (s *KonturService) StartKonturGame(n int) ([]*ent.KonturResponse, error) {
	kontur, err := s.repo.StartKonturGame(n)
	if err != nil {
		return nil, err
	}
	return kontur, nil
}
func (s *KonturService) ProcessKonturGame(params *ent.ProcessRequest, userId int64) (*ent.ProcessResponse, error) {
	kontur, err := s.repo.ProcessKonturGame(params, userId)
	if err != nil {
		return nil, err
	}
	return kontur, nil
}
