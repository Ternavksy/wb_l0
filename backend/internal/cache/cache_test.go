package cache

import (
	"context"
	"fmt"
	"reflect"
	"sync"
	"testing"
	"time"

	"wb_l0/internal/db"
	"wb_l0/internal/models"

	"go.uber.org/zap"
)

func TestCache(t *testing.T) {
	logger = zap.NewNop()

	// Очищаем кэш перед тестом
	cache.Lock()
	cache.m = make(map[string]models.Order)
	cache.Unlock()

	log, err := zap.NewDevelopment()
	if err != nil {
		t.Fatalf("failed to initialize logger: %v", err)
	}
	if err := db.InitDB(log); err != nil {
		t.Fatalf("failed to initialize db: %v", err)
	}
	defer db.Pool().Close()

	t.Run("AddToCacheAndGetFromCache", func(t *testing.T) {
		order := models.Order{
			OrderUID:    "test-order-1",
			TrackNumber: "TRACK123",
			Entry:       "TEST",
			CustomerID:  "test_customer",
			DateCreated: time.Now(),
			Locale:      "en",
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
				Transaction: "test_transaction",
				Currency:    "USD",
				Provider:    "test_provider",
				Amount:      1500,
			},
			Items: []models.Item{
				{
					ChrtID:      123456,
					TrackNumber: "TRACK123",
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

		AddToCache(order)

		retrievedOrder, exists := GetFromCache("test-order-1")
		if !exists {
			t.Error("expected order to exist in cache")
		}
		if !reflect.DeepEqual(order, *retrievedOrder) {
			t.Errorf("retrieved order does not match saved order.\nGot: %+v\nWant: %+v", *retrievedOrder, order)
		}

		_, exists = GetFromCache("non-existent-order")
		if exists {
			t.Error("expected non-existent order to return false")
		}
	})

	t.Run("InitCache", func(t *testing.T) {
		// Очищаем кэш
		cache.Lock()
		cache.m = make(map[string]models.Order)
		cache.Unlock()

		// Добавляем тестовый заказ в базу
		testOrder := models.Order{
			OrderUID:    "test-order-init",
			TrackNumber: "INIT123",
			Entry:       "TEST",
			CustomerID:  "test_customer",
			DateCreated: time.Now(),
			Locale:      "en",
		}
		err := db.SaveOrder(context.Background(), testOrder)
		if err != nil {
			t.Fatalf("failed to save order for InitCache: %v", err)
		}
		defer func() {
			_, _ = db.Pool().Exec(context.Background(), "DELETE FROM orders WHERE order_uid = $1", testOrder.OrderUID)
		}()

		// Инициализируем кэш
		err = InitCache(log)
		if err != nil {
			t.Fatalf("failed to initialize cache: %v", err)
		}

		// Проверяем что заказ загружен в кэш
		_, exists := GetFromCache(testOrder.OrderUID)
		if !exists {
			t.Error("expected order to be loaded into cache")
		}
	})

	t.Run("ConcurrentAccess", func(t *testing.T) {
		// Очищаем кэш
		cache.Lock()
		cache.m = make(map[string]models.Order)
		cache.Unlock()

		var wg sync.WaitGroup
		numGoroutines := 100
		wg.Add(numGoroutines * 2) // Для чтения и записи

		// Записываем заказы в кэш
		for i := 0; i < numGoroutines; i++ {
			go func(id int) {
				defer wg.Done()
				order := models.Order{OrderUID: fmt.Sprintf("order-%d", id)}
				AddToCache(order)
			}(i)
		}

		// Читаем из кэша
		for i := 0; i < numGoroutines; i++ {
			go func(id int) {
				defer wg.Done()
				_, _ = GetFromCache(fmt.Sprintf("order-%d", id))
			}(i)
		}

		wg.Wait()

		// Проверяем, что все заказы записаны
		for i := 0; i < numGoroutines; i++ {
			_, exists := GetFromCache(fmt.Sprintf("order-%d", i))
			if !exists {
				t.Errorf("expected order-%d to exist in cache", i)
			}
		}
	})
}
