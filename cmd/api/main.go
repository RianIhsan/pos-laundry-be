package main

import (
	"log"
	"os"

	"github.com/RianIhsan/pos-laundry-be/config"
	"github.com/RianIhsan/pos-laundry-be/internal/server"
	pg "github.com/RianIhsan/pos-laundry-be/pkg/db"
	"github.com/RianIhsan/pos-laundry-be/pkg/logger"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.NewAppConfig(os.Getenv("config"))
	if err != nil {
		log.Fatal(err)
	}

	apiLogger := logger.NewLogrusLogger(cfg)

	psqlDB, err := pg.NewPostgresConnection(cfg)
	if err != nil {
		apiLogger.Fatalf("init postgres connection failed: %v", err)
	}
	apiLogger.Info("init postgres connection success")

	err = pg.Migrate(psqlDB)
	if err != nil {
		apiLogger.Fatalf("migrate failed: %v", err)
	}
	apiLogger.Info("migrate success")

	app := gin.New()
	err = app.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		apiLogger.Fatalf("SetTrustedProxies failed: %v", err)
	}
	app.Use(gin.Recovery())

	// Start
	s := server.NewServer(&server.ServerConfig{
		App:    app,
		Logger: apiLogger,
		Cfg:    cfg,
		Db:     psqlDB,
	})

	if err := s.Run(); err != nil {
		apiLogger.Fatalf("server run failed: %v", err)
	}

}
