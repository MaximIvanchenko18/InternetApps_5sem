package app

import (
	"bytes"
	"fmt"
	"net/http"
	"path/filepath"

	"mime/multipart"

	"InternetApps_5sem/internal/app/role"

	"crypto/sha1"
	"encoding/hex"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
)

func (app *Application) uploadImage(c *gin.Context, image *multipart.FileHeader, UUID string) (*string, error) {
	src, err := image.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	extension := filepath.Ext(image.Filename)
	if extension != ".jpg" && extension != ".jpeg" {
		return nil, fmt.Errorf("разрешены только изображения формата jpeg")
	}
	imageName := UUID + extension

	_, err = app.minioClient.PutObject(c, app.config.Minio.BucketName, imageName, src, image.Size, minio.PutObjectOptions{
		ContentType: "image/jpeg",
	})
	if err != nil {
		return nil, err
	}
	imageURL := fmt.Sprintf("http://%s/%s/%s", app.config.Minio.Endpoint, app.config.Minio.BucketName, imageName)
	return &imageURL, nil
}

func (app *Application) deleteImage(c *gin.Context, UUID string) error {
	imageName := UUID + ".jpg"

	err := app.minioClient.RemoveObject(c, app.config.Minio.BucketName, imageName, minio.RemoveObjectOptions{})
	if err != nil {
		return err
	}

	return nil
}

func getUserId(c *gin.Context) string {
	userId, _ := c.Get("userId")
	return userId.(string)
}

func getUserRole(c *gin.Context) role.Role {
	userRole, _ := c.Get("userRole")
	return userRole.(role.Role)
}

func generateHashString(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func shipmentRequest(flight_id string) error {
	url := "http://localhost:4000/api/shipment"
	payload := fmt.Sprintf(`{"flight_id": "%s"}`, flight_id)

	resp, err := http.Post(url, "application/json", bytes.NewBufferString(payload))
	if err != nil {
		return err
	}

	if resp.StatusCode >= 300 {
		return fmt.Errorf(`shipment failed with status: {%s}`, resp.Status)
	}

	return nil
}
