package cache

import (
	"context"
	"sync"
	"wb_l0/internal/db"
	"wb_l0/internal/models"

	"go.uber.org/zap"
)

var cache = struct {
	sync.RWMutex
	m map[string]models.Order
}{m: make(map[string]models.Order)}

var logger *zap.Logger

func InitCache(l *zap.Logger) error {
	logger = l
	rows, err := db.Pool().Query(context.Background(), `SELECT order_uid FROM orders`)
	if err != nil {
		logger.Error("Failed to load orders for cache", zap.Error(err))
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var uid string
		if err := rows.Scan(&uid); err != nil {
			logger.Error("Failed to scan order_uid", zap.Error(err))
			return err
		}
		order, err := db.GetOrder(context.Background(), uid)
		if err != nil {
			logger.Error("Failed to get order for cache", zap.Error(err), zap.String("order_uid", uid))
			continue
		}
		cache.Lock()
		cache.m[uid] = *order
		cache.Unlock()
		logger.Info("Loaded order into cache", zap.String("order_uid", uid))
	}
	logger.Info("Cache initialized", zap.Int("order_count", len(cache.m)))
	return nil
}

func GetFromCache(orderUID string) (*models.Order, bool) {
	cache.RLock()
	order, exists := cache.m[orderUID]
	cache.RUnlock()
	if exists {
		logger.Info("Cache hit", zap.String("order_uid", orderUID))
	} else {
		logger.Info("Cache miss", zap.String("order_uid", orderUID))
	}
	return &order, exists
}

func AddToCache(order models.Order) {
	cache.Lock()
	cache.m[order.OrderUID] = order
	cache.Unlock()
	logger.Info("Added to cache", zap.String("order_uid", order.OrderUID))
}
