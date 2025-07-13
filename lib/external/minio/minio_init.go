package minio

import (
	"backend/app/config"
	"backend/app/database"
	"backend/pkg/models"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var minioClient *minio.Client

func MinioInit() {
	var err error
	minioClient, err = minio.New(config.AppConfig.MinioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AppConfig.MinioAccessKey, config.AppConfig.MinioSecretKey, ""),
		Secure: false, // Set to true if MinIO is using HTTPS
	})
	if err != nil {
		log.Fatalf("Failed to initialize MinIO client: %v", err)
	}
}

// UploadFile godoc
// @Summary Upload a file to MinIO
// @Description Upload a profile picture to MinIO
// @Tags Upload
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "Profile Picture"
// @Success 200 {object} map[string]string "File successfully uploaded"
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /api/upload [post]
func UploadFile(c echo.Context) error {
	// Get the file from the form-data
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "No file is received"})
	}

	userId := c.FormValue("id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "user id"})
	}

	// Open the file
	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to open the file"})
	}
	defer src.Close()

	// Generate file name (or you can use file.Filename)
	objectName := fmt.Sprintf("uploads/%s", file.Filename)

	contentType := "image/jpeg"
	// Upload the file to MinIO
	_, err = minioClient.PutObject(c.Request().Context(), config.AppConfig.MinioBuckect, objectName, src, file.Size, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": fmt.Sprintf("Failed to upload file: %v", err)})
	}

	// Menghasilkan URL file
	url := fmt.Sprintf("http://%s/%s/%s", config.AppConfig.MinioEndpoint, config.AppConfig.MinioBuckect, objectName)
	fmt.Println("url", url)
	db := database.InitPostgresGorm()
	var user models.User
	if err := db.Where("id=?", userId).First(&user).Error; err != nil {
		return err
	}

	user.Profile = url

	if err := db.Save(&user).Error; err != nil {
		return err
	}

	// Return response
	return c.JSON(http.StatusOK, map[string]string{
		"message": "File successfully uploaded",
		"file":    objectName,
	})
}
