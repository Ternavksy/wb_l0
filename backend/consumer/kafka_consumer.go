package consumer

import (
	"context"
	"encoding/json"
	"time"
	"wb_l0/internal/cache"
	"wb_l0/internal/db"
	"wb_l0/internal/models"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

var logger *zap.Logger

func StartConsumer(l *zap.Logger) {
	logger = l
	for {
		reader := kafka.NewReader(kafka.ReaderConfig{
			Brokers:   []string{"localhost:9092"},
			Topic:     "orders",
			Partition: 0,
			MinBytes:  10e3,
			MaxBytes:  10e6,
			MaxWait:   1 * time.Second,
		})

		logger.Info("Kafka consumer started", zap.String("topic", "orders"))

		for {
			msg, err := reader.ReadMessage(context.Background())
			if err != nil {
				logger.Error("Failed to read message from Kafka", zap.Error(err))
				reader.Close()
				time.Sleep(5 * time.Second)
				break
			}
			var order models.Order
			if err := json.Unmarshal(msg.Value, &order); err != nil {
				logger.Error("Failed to unmarshal order JSON", zap.Error(err), zap.ByteString("message", msg.Value))
				continue
			}
			if order.OrderUID == "" {
				logger.Error("Received order with empty order_uid", zap.ByteString("message", msg.Value))
				continue
			}

			if err := db.SaveOrder(context.Background(), order); err != nil {
				logger.Error("Failed to save order to database", zap.Error(err), zap.String("order_uid", order.OrderUID))
				continue
			}

			cache.AddToCache(order)
			logger.Info("Processed order", zap.String("order_uid", order.OrderUID))
		}
	}
}
