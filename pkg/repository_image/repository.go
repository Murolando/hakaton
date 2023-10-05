package repositoryImage

import (
	"image"
	"mime/multipart"

	repoS3 "github.com/Murolando/hakaton_geo/pkg/repository_image/repo_s3"
)

type RepoS3 interface {
	UpdateUserImage(file *multipart.FileHeader, userId int64) (string, error)
	DeleteUserImage(imageName string) error
	ResizeImage(file *multipart.FileHeader) (image.Image, error)
	UniquePart(id int64) (*string, error)
}

type Image struct {
	RepoS3
}

func NewImage() *Image {
	return &Image{
		RepoS3: repoS3.NewS3WebStorage(),
	}
}
