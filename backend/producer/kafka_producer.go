package main

import (
	"context"
	"encoding/json"
	"time"
	"wb_l0/internal/models"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}
	defer logger.Sync()

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "orders",
		Balancer: &kafka.LeastBytes{},
	})
	defer writer.Close()

	order := models.Order{
		OrderUID:    "test-order-145",
		TrackNumber: "TESTTRACK145",
		Entry:       "NEWTEST145",
		Delivery: models.Delivery{
			Name:    "Zoe Doe",
			Phone:   "+880055535123",
			Zip:     "1111",
			City:    "Test New",
			Address: "111 Test St",
			Region:  "New Test Region",
			Email:   "Newtest@example.com",
		},
		Payment: models.Payment{
			Transaction:  "test-trans-145",
			Currency:     "RUS",
			Provider:     "newtestpay",
			Amount:       777,
			PaymentDt:    time.Now().Unix(),
			Bank:         "WBBank",
			DeliveryCost: 777,
			GoodsTotal:   111,
			CustomFee:    0,
		},
		Items: []models.Item{
			{
				ChrtID:      12345,
				TrackNumber: "TESTTRACK145",
				Price:       500,
				Rid:         "test-rid-145",
				Name:        "New Test Item 145",
				Sale:        90,
				Size:        "S",
				TotalPrice:  450,
				NmID:        67890,
				Brand:       "New Test Brand",
				Status:      100,
			},
		},
		Locale:            "ru",
		InternalSignature: nil,
		CustomerID:        "new-test-customer-145",
		DeliveryService:   "new-test-delivery-145",
		ShardKey:          "2",
		SMID:              1,
		DateCreated:       time.Now(),
		OOFShard:          "1",
	}

	msg, err := json.Marshal(order)
	if err != nil {
		logger.Error("Failed to marshal order", zap.Error(err))
		return
	}
	err = writer.WriteMessages(context.Background(),
		kafka.Message{
			Value: msg,
		},
	)
	if err != nil {
		logger.Error("Failed to write message to Kafka", zap.Error(err))
		return
	}
	logger.Info("Test order sent", zap.String("order_uid", order.OrderUID))
}
