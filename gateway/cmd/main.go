package main

import (
	"log"

	"txn-engine-phase-2/gateway/internal/client"
	"txn-engine-phase-2/gateway/internal/config"
	"txn-engine-phase-2/gateway/internal/handler"
	"txn-engine-phase-2/gateway/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	adminClient := client.NewAdminClient(cfg.AdminServiceURL)

	authHandler := handler.NewAuthHandler(adminClient)

	r := gin.Default()
	router.RegisterRoutes(r, authHandler)

	log.Printf("Gateway running on port %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
