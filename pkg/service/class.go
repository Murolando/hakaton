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

func (s *ClassService) DashboardClass(userId int64) (*ent.ChildDashClassResponce, error) {
	var responce ent.ChildDashClassResponce
	class,err := s.repo.DashboardClass(userId)
	if err!=nil{
		return nil,err
	}
	responce.ClassProgress = class
	
	r,err:= s.repo.CommonProgressInfo(userId)
	if err !=nil{
		return nil,err
	}
	responce.MaxExProgressBar = r.MaxExProgressBar
	responce.MaxTheoryProgressBar = r.MaxTheoryProgressBar
	responce.TheoryProgressBar = r.TheoryProgressBar
	responce.ExProgressBar = r.ExProgressBar
	return &responce,nil
}


func (s *ClassService) MyClass(userId int64)([]*ent.ChildMyClassResponce,error){
	return s.repo.Class.MyClass(userId)
}