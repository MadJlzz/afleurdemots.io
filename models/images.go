package models

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Image is NOT stored in the database
type Image struct {
	GalleryID uint
	Filename string
}

func (i *Image) RelativePath() string {
	return fmt.Sprintf("images/galleries/%v/%v", i.GalleryID, i.Filename)
}

func (i *Image) Path() string {
	return "/" + i.RelativePath()
}

type ImageService interface {
	Create(galleryID uint, filename string, r io.ReadCloser) error
	ByGalleryID(galleryID uint) ([]Image, error)
	Delete(image *Image) error
}

func NewImageService() ImageService {
	return &imageService{}
}

type imageService struct {}

func (is *imageService) Delete(image *Image) error {
	return os.Remove(image.RelativePath())
}

func (is *imageService) Create(galleryID uint, filename string, r io.ReadCloser) error {
	defer r.Close()
	path, err := is.mkImagePath(galleryID)
	if err != nil {
		return err
	}

	dst, err := os.Create(path + filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, r)
	if err != nil {
		return err
	}

	return nil
}

func (is *imageService) ByGalleryID(galleryID uint) ([]Image, error) {
	path := is.imagePath(galleryID)
	imgStrings, err := filepath.Glob(path + "*")
	if err != nil {
		return nil, err
	}
	ret := make([]Image, len(imgStrings))
	for i := range imgStrings {
		imgStrings[i] = strings.Replace(imgStrings[i], path, "", 1)
		ret[i] = Image{
			GalleryID: galleryID,
			Filename:  imgStrings[i],
		}
	}
	return ret, err
}

func (is *imageService) imagePath(galleryID uint) string {
	return fmt.Sprintf("images/galleries/%v/", galleryID)
}

func (is *imageService) mkImagePath(galleryID uint) (string, error) {
	galleryPath := is.imagePath(galleryID)
	err := os.MkdirAll(galleryPath, 0755)
	if err != nil {
		return "", nil
	}
	return galleryPath, nil
}