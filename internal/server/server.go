package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/RianIhsan/pos-laundry-be/config"
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const (
	ctxTimeout = 5 * time.Second
)

type ServerConfig struct {
	App    *gin.Engine
	Logger *logrus.Logger
	Cfg    *config.Config
	Db     *gorm.DB
}

type Server struct {
	app    *gin.Engine
	logger *logrus.Logger
	cfg    *config.Config
	db     *gorm.DB
}

func NewServer(config *ServerConfig) *Server {
	return &Server{
		app:    config.App,
		logger: config.Logger,
		cfg:    config.Cfg,
		db:     config.Db,
	}
}

func (s *Server) Run() error {
	if err := s.initServer(); err != nil {
		return err
	}

	apiServer := s.newApiServer()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	serverError := make(chan error, 1)
	s.startServer(serverError, apiServer)

	if err := s.waitForShutdown(ctx, serverError); err != nil {
		return err
	}

	s.gracefulShutdown(apiServer)
	s.cleanUp()

	s.logger.Info("Server exited properly")
	return nil
}

func (s *Server) newApiServer() *http.Server {
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			"GET", "POST", "PUT", "DELETE", "OPTIONS",
		},
		AllowedHeaders: []string{
			"Content-Type",
			"Authorization",
		},
		AllowCredentials: true,
	})
	return &http.Server{
		Addr:         fmt.Sprintf("%s:%d", s.cfg.Server.Host, s.cfg.Server.Port),
		ReadTimeout:  time.Second * s.cfg.Server.ReadTimeout,
		WriteTimeout: time.Second * s.cfg.Server.WriteTimeout,
		Handler:      corsMiddleware.Handler(s.app),
	}
}

func (s *Server) initServer() error {
	if err := s.Bootstrap(); err != nil {
		s.logger.Fatalf("Failed to bootstrap the server: %v", err)
	}
	s.logger.Info("Server bootstrapped successfully")
	return nil
}

func (s *Server) startServer(ch chan error, servers ...*http.Server) {
	for _, srv := range servers {
		go func(srv *http.Server) {
			s.logger.Infof("Server listening on %s", srv.Addr)
			ch <- srv.ListenAndServe()
		}(srv)
	}
}

func (s *Server) waitForShutdown(ctx context.Context, ch chan error) error {
	select {
	case err := <-ch:
		if err != nil && err != http.ErrServerClosed {
			return fmt.Errorf("server error: %w", err)
		}
	case <-ctx.Done():
		s.logger.Info("Shutdown signal received")
	}
	return nil
}

func (s *Server) gracefulShutdown(servers ...*http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), ctxTimeout)
	defer cancel()

	for _, srv := range servers {
		s.logger.Infof("Shutting down server on %s", srv.Addr)
		if err := srv.Shutdown(ctx); err != nil {
			s.logger.Errorf("Error shutting down server on %s: %v", srv.Addr, err)
		} else {
			s.logger.Infof("Server on %s shut down gracefully", srv.Addr)
		}
	}
}

func (s *Server) cleanUp() {
	if s.db != nil {
		sqlDB, err := s.db.DB()
		if err != nil {
			s.logger.Errorf("Error getting DB from GORM: %v", err)
			return
		}
		if err := sqlDB.Close(); err != nil {
			s.logger.Errorf("Error closing DB connection: %v", err)
		} else {
			s.logger.Info("Database connection closed")
		}
	}
}
