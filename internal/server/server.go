package server

import (
	"evermos_rakamin/internal/config"
	"evermos_rakamin/internal/common"
	"evermos_rakamin/internal/http/router"
	"evermos_rakamin/internal/util"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"
	"net/http"
    "errors"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

type Server struct {
	app *fiber.App
	cfg *config.Config
}

func NewServer(cfg *config.Config, publicRoutes, privateRoutes []*router.Route) *Server {
	app := fiber.New()

	// Register Public Routes
	for _, route := range publicRoutes {
		app.Add(route.Method, route.Path, route.Handler)
	}

	// Register Private Routes with JWT Middleware
	for _, v := range privateRoutes {
		if v.Roles[0] == "admin" && len(v.Roles) == 1{
			app.Add(v.Method, v.Path, JWTMiddleware(cfg.JWTSecretKey), 
				func(ctx *fiber.Ctx) error {
					err := CheckIsAdmin(ctx, cfg.JWTSecretKey)
					if err != nil {
						return util.JSONResponse(ctx, http.StatusUnauthorized, "Unauthorized", err, nil)
					}
					return ctx.Next()
				}, v.Handler)
		} else {
			app.Add(v.Method, v.Path, JWTMiddleware(cfg.JWTSecretKey), v.Handler)
		}
	}
	return &Server{app, cfg}
}

func (s *Server) Run() {
	go func() {
		if err := s.app.Listen(fmt.Sprintf(":%s", s.cfg.Port)); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()
}

func (s *Server) GracefulShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	fmt.Println("Shutting down server...")

	if err := s.app.Shutdown(); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	<-ctx.Done()
	fmt.Println("Server exited gracefully")
}


func CheckIsAdmin(ctx *fiber.Ctx, secretKey string) error {
	token := ctx.Get("token")

	if token == "" {
		return errors.New("token empty")
	}

	claims, err := util.DecodeJWT(token, secretKey)
	if err != nil {
		return errors.New("unauthenticated")
	}

	isAdmin := claims["is_admin"]
	if isAdmin == false {
		return errors.New("Not Admin")
	}
	return nil
}

func JWTMiddleware(secretKey string) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(secretKey),
		ContextKey:   "user",
		Claims:       &common.JwtCustomClaims{},
		ErrorHandler: jwtError,
		TokenLookup: "header:token",
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"success": false,
		"message": "Unauthorized",
		"error":   err.Error(),
		"data":    nil,
	})
}