package main

import (
	"evermos_rakamin/internal/config"
	"evermos_rakamin/internal/entity"
	"evermos_rakamin/internal/builder"
	"evermos_rakamin/internal/common"
	"evermos_rakamin/internal/database"
	"evermos_rakamin/internal/server"

	"fmt"
	"os"
	"regexp"
)

func main() {
	projectDirName := "Golang-API-Mini-Project-Evermos-Rakamin"
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	cfg, err := config.NewConfig(string(rootPath) + "/.env")
	if err != nil {
		panic(err)
	}

	db, err := database.ConnectToMysql(cfg)
	if err != nil {
		panic(err)
	}

	common.AppConfig = cfg

	publicRoutes := builder.BuildPublicRoutes(db, cfg)
	privateRoute := builder.BuildPrivateRoutes(db, cfg)

	err = db.AutoMigrate(
		&entity.User{},
		&entity.Alamat{},
		&entity.Toko{},
		&entity.Category{},
		&entity.Produk{},
		&entity.FotoProduk{},
		&entity.LogProduk{},
		&entity.Trx{},
		&entity.DetailTrx{},
	)
	if err != nil {
		panic("Failed to migrate database")
	}

	fmt.Println("Database migration success!")

	srv := server.NewServer(cfg, publicRoutes, privateRoute)
	srv.Run()
	srv.GracefulShutdown()
}
