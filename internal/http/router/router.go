package router

import (
	"evermos_rakamin/internal/http/handler"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

const (
	Admin = "admin"
	User  = "user"
)

var (
	allRoles  = []string{Admin, User}
	onlyAdmin = []string{Admin}
	onlyUser  = []string{User}
)

type Route struct {
	Method  string
	Path    string
	Handler fiber.Handler
	Roles   []string
}

func PublicRoutes(AuthHandler *handler.AuthHandler, UserHandler *handler.UserHandler) []*Route {
	return []*Route{
		{Method: http.MethodPost, Path: "/api/v1/auth/login", Handler: AuthHandler.Login, Roles: allRoles},
		{Method: http.MethodPost, Path: "/api/v1/auth/register", Handler: AuthHandler.Register, Roles: allRoles},
		{Method: http.MethodPost, Path: "/user/generate-password", Handler: UserHandler.GeneratePassword, Roles: onlyUser},
		{Method: http.MethodPut, Path: "/api/v1/user/:id", Handler: UserHandler.UpdateUser, Roles: onlyAdmin},
	}
}

func PrivateRoutes(UserHandler *handler.UserHandler, AlamatHandler *handler.AlamatHandler, TokoHandler *handler.TokoHandler, CategoryHandler *handler.CategoryHandler, ProdukHandler *handler.ProdukHandler) []*Route {	return []*Route{
		{Method: http.MethodGet, Path: "/api/v1/user", Handler: UserHandler.GetProfile, Roles: allRoles},
		{Method: http.MethodPut, Path: "/api/v1/user", Handler: UserHandler.UpdateProfile, Roles: allRoles},
		{Method: http.MethodGet, Path: "/api/v1/all-user", Handler: UserHandler.GetAllUser, Roles: allRoles},

		// Alamat Routes
		{Method: http.MethodGet, Path: "/api/v1/user/alamat", Handler: AlamatHandler.GetAlamatUser, Roles: allRoles},
		{Method: http.MethodGet, Path: "/api/v1/user/alamat/:id", Handler: AlamatHandler.GetAlamatById, Roles: allRoles},
		{Method: http.MethodPost, Path: "/api/v1/user/alamat", Handler: AlamatHandler.CreateAlamat, Roles: allRoles},
		{Method: http.MethodPut, Path: "/api/v1/user/alamat/:id", Handler: AlamatHandler.UpdateAlamat, Roles: allRoles},
		{Method: http.MethodDelete, Path: "/api/v1/user/alamat/:id", Handler: AlamatHandler.DeleteAlamat, Roles: allRoles},

		// Toko Routes
		{Method: http.MethodGet, Path: "/api/v1/toko/my", Handler: TokoHandler.GetMyToko, Roles: allRoles},
		{Method: http.MethodGet, Path: "/api/v1/toko/:id_toko", Handler: TokoHandler.GetTokoByID, Roles: allRoles},
		{Method: http.MethodGet, Path: "/api/v1/toko", Handler: TokoHandler.GetTokoPaginated, Roles: allRoles},
		{Method: http.MethodPut, Path: "/api/v1/toko/:id_toko", Handler: TokoHandler.UpdateToko, Roles: allRoles},
		
		// Category Routes
		{Method: http.MethodGet, Path: "/api/v1/category", Handler: CategoryHandler.GetAllCategories, Roles: allRoles},
		{Method: http.MethodGet, Path: "/api/v1/category/:id", Handler: CategoryHandler.GetCategoryByID, Roles: allRoles},
		{Method: http.MethodPost, Path: "/api/v1/category", Handler: CategoryHandler.CreateCategory, Roles: onlyAdmin},
		{Method: http.MethodPut, Path: "/api/v1/category/:id", Handler: CategoryHandler.UpdateCategory, Roles: onlyAdmin},
		{Method: http.MethodDelete, Path: "/api/v1/category/:id", Handler: CategoryHandler.DeleteCategory, Roles: onlyAdmin},

		// Produk Routes
		{Method: http.MethodGet, Path: "/api/v1/product", Handler: ProdukHandler.GetAllProduk, Roles: allRoles},
		{Method: http.MethodGet, Path: "/api/v1/product/:id", Handler: ProdukHandler.GetProdukByID, Roles: allRoles},
		{Method: http.MethodPost, Path: "/api/v1/product", Handler: ProdukHandler.CreateProduk, Roles: onlyUser},
		{Method: http.MethodPut, Path: "/api/v1/product/:id", Handler: ProdukHandler.UpdateProduk, Roles: onlyUser},
		{Method: http.MethodDelete, Path: "/api/v1/product/:id", Handler: ProdukHandler.DeleteProduk, Roles: onlyUser},
	}
}

