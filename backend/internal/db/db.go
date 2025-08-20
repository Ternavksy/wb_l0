package db

import (
	"context"
	"wb_l0/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

var pool *pgxpool.Pool
var logger *zap.Logger

func InitDB(l *zap.Logger) error {
	logger = l
	connString := "postgres://postgres:postgres@localhost:5432/orders_db?sslmode=disable"
	var err error
	pool, err = pgxpool.New(context.Background(), connString)
	if err != nil {
		logger.Error("Failed to connect to disable", zap.Error(err))
		return err
	}
	logger.Info("Connected to database")
	return nil
}

func SaveOrder(ctx context.Context, order models.Order) error {
	tx, err := pool.Begin(ctx)
	if err != nil {
		logger.Error("Failed to start transaction", zap.Error(err))
		return err
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx,
		`INSERT INTO orders (order_uid, track_number, entry, customer_id, date_created, locale, internal_signature, delivery_service, shardkey, sm_id, oof_shard)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
		order.OrderUID, order.TrackNumber, order.Entry, order.CustomerID, order.DateCreated, order.Locale, order.InternalSignature, order.DeliveryService, order.ShardKey, order.SMID, order.OOFShard)
	if err != nil {
		logger.Error("Failed to insert into orders", zap.Error(err))
		return err
	}

	_, err = tx.Exec(ctx,
		`INSERT INTO delivery (order_uid, name, phone, zip, city, address, region, email)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		order.OrderUID, order.Delivery.Name, order.Delivery.Phone, order.Delivery.Zip, order.Delivery.City, order.Delivery.Address, order.Delivery.Region, order.Delivery.Email)
	if err != nil {
		logger.Error("Failed to insert into delivery", zap.Error(err))
		return err
	}

	_, err = tx.Exec(ctx,
		`INSERT INTO payment (order_uid, transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
		order.OrderUID, order.Payment.Transaction, order.Payment.RequestID, order.Payment.Currency, order.Payment.Provider, order.Payment.Amount, order.Payment.PaymentDt, order.Payment.Bank, order.Payment.DeliveryCost, order.Payment.GoodsTotal, order.Payment.CustomFee)
	if err != nil {
		logger.Error("Failed to insert into payment", zap.Error(err))
		return err
	}
	for _, item := range order.Items {
		_, err = tx.Exec(ctx,
			`INSERT INTO items (order_uid, chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status)
            VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`,
			order.OrderUID, item.ChrtID, item.TrackNumber, item.Price, item.Rid, item.Name, item.Sale, item.Size, item.TotalPrice, item.NmID, item.Brand, item.Status)
		if err != nil {
			logger.Error("Failed to insert into items", zap.Error(err))
			return err
		}
	}
	err = tx.Commit(ctx)
	if err != nil {
		logger.Error("Failed to commit transaction", zap.Error(err))
		return err
	}
	logger.Info("Order saved", zap.String("order_uid", order.OrderUID))
	return nil
}

func GetOrder(ctx context.Context, orderUID string) (*models.Order, error) {
	order := &models.Order{OrderUID: orderUID}
	// Чтение из orders
	err := pool.QueryRow(ctx,
		`SELECT track_number, entry, customer_id, date_created, locale, internal_signature, delivery_service, shardkey, sm_id, oof_shard
        FROM orders WHERE order_uid = $1`,
		orderUID).Scan(&order.TrackNumber, &order.Entry, &order.CustomerID, &order.DateCreated, &order.Locale, &order.InternalSignature, &order.DeliveryService, &order.ShardKey, &order.SMID, &order.OOFShard)
	if err != nil {
		logger.Error("Failed to get order", zap.Error(err))
		return nil, err
	}
	// Чтение из delivery
	err = pool.QueryRow(ctx,
		`SELECT name, phone, zip, city, address, region, email
        FROM delivery WHERE order_uid = $1`, orderUID).Scan(&order.Delivery.Name, &order.Delivery.Phone, &order.Delivery.Zip, &order.Delivery.City, &order.Delivery.Address, &order.Delivery.Region, &order.Delivery.Email)
	if err != nil {
		logger.Error("Failed to get delivery", zap.Error(err))
		return nil, err
	}
	// Чтение payment
	err = pool.QueryRow(ctx, `
        SELECT transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee
        FROM payment WHERE order_uid = $1`,
		orderUID).Scan(&order.Payment.Transaction, &order.Payment.RequestID, &order.Payment.Currency, &order.Payment.Provider, &order.Payment.Amount, &order.Payment.PaymentDt, &order.Payment.Bank, &order.Payment.DeliveryCost, &order.Payment.GoodsTotal, &order.Payment.CustomFee)
	if err != nil {
		logger.Error("Failed to get payment", zap.Error(err))
		return nil, err
	}
	rows, err := pool.Query(ctx,
		`SELECT chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status
        FROM items WHERE order_uid = $1`, orderUID)
	if err != nil {
		logger.Error("Failed to get items", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item models.Item
		err = rows.Scan(&item.ChrtID, &item.TrackNumber, &item.Price, &item.Rid, &item.Name, &item.Sale, &item.Size, &item.TotalPrice, &item.NmID, &item.Brand, &item.Status)
		if err != nil {
			logger.Error("Failed to scan item", zap.Error(err))
			return nil, err
		}
		order.Items = append(order.Items, item)
	}
	return order, nil
}

func Pool() *pgxpool.Pool {
	return pool
}
