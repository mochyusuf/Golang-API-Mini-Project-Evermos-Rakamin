package util

import (
	"evermos_rakamin/internal/helper"
	"mime/multipart"

	"fmt"
	"path/filepath"

	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

var imageDir = "upload/image"
var fileUploadDir = filepath.Join(helper.ProjectDirectory, imageDir)

func UploadFileToko(ctx *fiber.Ctx, file *multipart.FileHeader, fileName string) (string, error) {
	NewfileName := replaceFileName(fileName)
	if err := ctx.SaveFile(file, fmt.Sprintf("./%s/toko/%s", imageDir, NewfileName)); err != nil {
		return "", err
	}

	return fileName, nil
}

func UploadFileProduk(ctx *fiber.Ctx, file *multipart.FileHeader, fileName string) (string, error) {
	NewfileName := replaceFileName(fileName)
	if err := ctx.SaveFile(file, fmt.Sprintf("./%s/produk/%s", imageDir, NewfileName)); err != nil {
		return "", err
	}

	return fileName, nil
}

func replaceFileName(filename string) string {
	newName := strconv.FormatInt(time.Now().Unix(), 10)
	return newName + "-" + filename
}
