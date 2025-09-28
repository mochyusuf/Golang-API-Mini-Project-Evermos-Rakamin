package builder

import (
	"evermos_rakamin/internal/config"
	"evermos_rakamin/internal/http/handler"
	"evermos_rakamin/internal/http/router"
	"evermos_rakamin/internal/repository"
	"evermos_rakamin/internal/service"

	"gorm.io/gorm"
)

func BuildPrivateRoutes(db *gorm.DB, cfg *config.Config) []*router.Route {
	fotoRepository := repository.NewFotoProdukRepository(db)

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(cfg, userRepository)
	userHandler := handler.NewUserHandler(userService)

	alamatRepository := repository.NewAlamatRepository(db)
	alamatService := service.NewAlamatService(cfg, alamatRepository)
	alamatHandler := handler.NewAlamatHandler(alamatService)

	tokoRepository := repository.NewTokoRepository(db)
	tokoService := service.NewTokoService(cfg, tokoRepository)
	tokoHandler := handler.NewTokoHandler(tokoService)

	categoryRepository := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepository)
	categoryHandler := handler.NewCategoryHandler(categoryService)
	
	produkRepository := repository.NewProdukRepository(db)
	produkService := service.NewProdukService(produkRepository, fotoRepository)
	produkHandler := handler.NewProdukHandler(produkService)
	
	trxRepository := repository.NewTrxRepository(db)
	trxService := service.NewTrxService(cfg, trxRepository)
	trxHandler := handler.NewTrxHandler(trxService)	

	return router.PrivateRoutes(userHandler, alamatHandler, tokoHandler, categoryHandler, produkHandler, trxHandler)
}

func BuildPublicRoutes(db *gorm.DB, cfg *config.Config) []*router.Route {
	userRepository := repository.NewUserRepository(db)
	tokoRepository := repository.NewTokoRepository(db)

	// Initialize services
	
	authService := service.NewAuthService(cfg, userRepository, tokoRepository)
	userService := service.NewUserService(cfg, userRepository)


	// Initialize handlers

	authHandler := handler.NewAuthHandler(authService)
	userHandler := handler.NewUserHandler(userService)

	return router.PublicRoutes(authHandler, userHandler)
}
