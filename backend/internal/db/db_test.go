package db

import (
	"context"
	"reflect"
	"testing"
	"time"

	"wb_l0/internal/models"

	"go.uber.org/zap"
)

func TestSaveAndGetOrder(t *testing.T) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		t.Fatalf("failed to initialize logger: %v", err)
	}

	if err := InitDB(logger); err != nil {
		t.Fatalf("failed to initialize db: %v", err)
	}

	ctx := context.Background()

	testOrder := models.Order{
		OrderUID:          "test_order_uid_123",
		TrackNumber:       "TESTTRACK123",
		Entry:             "TEST",
		CustomerID:        "test_customer",
		DateCreated:       time.Now(),
		Locale:            "en",
		InternalSignature: nil,
		DeliveryService:   "test_service",
		ShardKey:          "1",
		SMID:              99,
		OOFShard:          "1",
		Delivery: models.Delivery{
			Name:    "Test Name",
			Phone:   "+1234567890",
			Zip:     "12345",
			City:    "Test City",
			Address: "Test Address",
			Region:  "Test Region",
			Email:   "test@example.com",
		},
		Payment: models.Payment{
			Transaction:  "test_transaction",
			RequestID:    "",
			Currency:     "USD",
			Provider:     "test_provider",
			Amount:       1500,
			PaymentDt:    time.Now().Unix(),
			Bank:         "test_bank",
			DeliveryCost: 200,
			GoodsTotal:   1300,
			CustomFee:    0,
		},
		Items: []models.Item{
			{
				ChrtID:      123456,
				TrackNumber: "TESTTRACK123",
				Price:       1000,
				Rid:         "test_rid",
				Name:        "Test Item",
				Sale:        10,
				Size:        "0",
				TotalPrice:  900,
				NmID:        789,
				Brand:       "Test Brand",
				Status:      200,
			},
		},
	}

	err = SaveOrder(ctx, testOrder)
	if err != nil {
		t.Fatalf("failed to save order: %v", err)
	}

	retrievedOrder, err := GetOrder(ctx, testOrder.OrderUID)
	if err != nil {
		t.Fatalf("failed to get order: %v", err)
	}

	testOrder.DateCreated = retrievedOrder.DateCreated
	testOrder.Payment.PaymentDt = retrievedOrder.Payment.PaymentDt

	if !reflect.DeepEqual(testOrder, *retrievedOrder) {
		t.Errorf("retrieved order does not match saved order.\nGot: %+v\nWant: %+v", *retrievedOrder, testOrder)
	}
}
