package main

import (
	"fmt"
	"wb_l0/consumer"
	"wb_l0/internal/cache"
	"wb_l0/internal/db"
	"wb_l0/internal/server"

	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}
	defer logger.Sync()

	if err := db.InitDB(logger); err != nil {
		logger.Fatal("Failed to initialize database", zap.Error(err))
	}

	if err := cache.InitCache(logger); err != nil {
		logger.Fatal("Failed to initialize cache", zap.Error(err))
	}

	go consumer.StartConsumer(logger)
	go server.StartServer(logger)

	logger.Info("Service started")
	fmt.Println("Service started")

	select {}
}
