package server

import (
	"evermos_rakamin/internal/config"
	"evermos_rakamin/internal/common"
	"evermos_rakamin/internal/http/router"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

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
		app.Add(v.Method, v.Path, JWTMiddleware(cfg.JWTSecretKey), v.Handler)
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

func JWTMiddleware(secretKey string) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(secretKey),
		ContextKey:   "user",
		Claims:       &common.JwtCustomClaims{},
		ErrorHandler: jwtError,
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