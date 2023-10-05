package repoS3

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/nfnt/resize"
)

const (
	avatarBaseDirPhoto = "user_avatars/"
)

type S3WebStorage struct {
	SecretKey  string
	BucketName string
	S3         *s3.S3
}

func NewS3WebStorage() *S3WebStorage {
	svc, err := createClient()
	if err != nil {
		return nil
	}

	return &S3WebStorage{
		BucketName: os.Getenv("BUCKET_NAME"),
		S3:         svc,
	}
}
func createClient() (*s3.S3, error) {
	//using amason sdk
	var creds *credentials.Credentials
	creds = credentials.NewEnvCredentials()
	// fmt.Println(os.Getenv("REGION"))
	sess, err := session.NewSession(&aws.Config{
		// Region:      aws.String("ru-1"),
		Region:      aws.String(os.Getenv("REGION")),
		Credentials: creds,
		// Endpoint: aws.String("https://s3.timeweb.com/"),
		// Endpoint: aws.String("http://minio:9000/"),
		Endpoint:         aws.String(os.Getenv("ENDPOINT")),
		S3ForcePathStyle: aws.Bool(true),
	})
	if err != nil {
		return nil, err
	}
	_, err = sess.Config.Credentials.Get()
	if err != nil {
		return nil, err
	}
	svc := s3.New(sess)
	return svc, nil
}
func (r *S3WebStorage) UpdateUserImage(file *multipart.FileHeader, userId int64) (string, error) {
	// filename
	uniquePart, err := r.UniquePart(userId)
	if err != nil {
		return "", err
	}
	imagePlaceName := avatarBaseDirPhoto + *uniquePart

	fn := filepath.Base(file.Filename)
	fileParts := strings.Split(fn, ".")
	if len(fileParts) != 2 {
		return "", errors.New("bad file name or file ext")
	}
	fileExt := fileParts[1]
	fileName := imagePlaceName + "." + fileExt

	//resize image
	img, err := r.ResizeImage(file)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	// create formfile and copy image.Image to formfile
	if err := jpeg.Encode(buf, img, nil); err != nil {
		return "", err
	}

	// save photo
	aclRights := s3.ObjectCannedACLPublicRead // права доступа, тут просто на чтение для всех
	_, err = r.S3.PutObject(&s3.PutObjectInput{
		Body:   bytes.NewReader(buf.Bytes()),
		Bucket: aws.String(r.BucketName),
		Key:    aws.String(fileName),
		ACL:    &aclRights,
	})
	if err != nil {
		return "", err
	}
	return fileName, nil
}
func (r *S3WebStorage) DeleteUserImage(imageName string) error {

	_, err := r.S3.DeleteObject(&s3.DeleteObjectInput{
		Bucket: &r.BucketName,
		Key:    aws.String(imageName),
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *S3WebStorage) UniquePart(id int64) (*string, error) {
	uniquePart := strconv.FormatInt(id, 10)
	hash := fmt.Sprintf("%x", sha256.Sum256([]byte(uniquePart)))
	return &hash, nil
}

func (r *S3WebStorage) ResizeImage(file *multipart.FileHeader) (image.Image, error) {
	src, err := file.Open()
	if err != nil {
		return nil, err
	}

	defer src.Close()
	img, _, err := image.Decode(src)
	if err != nil {
		return nil, err
	}

	img = resize.Thumbnail(300, 300, img, resize.MitchellNetravali)

	return img, err
}
