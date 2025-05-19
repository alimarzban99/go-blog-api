package filer

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/alimarzban99/go-blog-api/config"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type Filer struct{}

type FileResponse struct {
	FileName string `json:"file_name"`
	FilePath string `json:"file_path"`
	URL      string `json:"url"`
}

const MaxFileSize = 4 * 1024 * 1024 // 4 مگابایت

func NewFiler() *Filer {
	return &Filer{}
}

func (f *Filer) Uploader(file *multipart.FileHeader) (*FileResponse, error) {

	if file.Size > MaxFileSize {
		return nil, errors.New("file too large")
	}

	valid, ext := f.isAllowedFile(file)
	if !valid {
		return nil, errors.New("invalid file type")
	}

	datePath := time.Now().Format("2006/01/02")
	folderPath := filepath.Join("uploads", datePath)
	if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
		return nil, errors.New("failed to create directory")
	}

	filename := f.generateRandomFilename(ext)
	fullPath := filepath.Join(folderPath, filename)

	src, err := file.Open()
	if err != nil {
		return nil, errors.New("failed to open uploaded file")
	}
	defer src.Close()

	// فایل مقصد رو بساز
	dst, err := os.Create(fullPath)
	if err != nil {
		return nil, errors.New("failed to create file")
	}
	defer dst.Close()

	// کپی فایل
	if _, err := io.Copy(dst, src); err != nil {
		return nil, errors.New("failed to save file")
	}

	baseUrl := filepath.Join(config.Config.App.URL, "preview")

	relativePath := filepath.Join(datePath, filename)
	url := fmt.Sprintf("%s/%s", baseUrl, filepath.ToSlash(relativePath))

	res := FileResponse{
		FileName: filename,
		FilePath: filepath.Join(datePath, filename),
		URL:      url,
	}
	return &res, nil
}

func (f *Filer) isAllowedFile(fileUpload *multipart.FileHeader) (bool, string) {
	allowedMimeTypes := map[string]string{
		"image/jpeg": ".jpg",
		"image/png":  ".png",
	}

	file, err := fileUpload.Open()
	if err != nil {
		return false, ""
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			log.Println(err)
		}
	}(file)

	// خواندن چند بایت ابتدایی برای تشخیص MIME
	buf := make([]byte, 512)
	_, err = file.Read(buf)
	if err != nil {
		return false, ""
	}

	mimeType := http.DetectContentType(buf)
	ext, ok := allowedMimeTypes[mimeType]
	if !ok {
		return false, ""
	}

	return true, ext
}

func (f *Filer) generateRandomFilename(extension string) string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b) + extension
}
