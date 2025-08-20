package server

import (
	"encoding/json"
	"net/http"
	"wb_l0/internal/cache"
	"wb_l0/internal/db"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

var logger *zap.Logger

func StartServer(l *zap.Logger) error {
	logger = l
	router := mux.NewRouter()

	router.HandleFunc("/order/{id}", GetOrderHandler).Methods("GET")

	logger.Info("Starting HTTP server", zap.String("address", ":8080"))

	return http.ListenAndServe(":8080", router)
}

func GetOrderHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderUID := vars["id"]

	order, exists := cache.GetFromCache(orderUID)
	if !exists {
		var err error
		order, err = db.GetOrder(r.Context(), orderUID)
		if err != nil {
			logger.Error("Failed to get order from database", zap.String("order_uid", orderUID), zap.Error(err))
			http.Error(w, "Order not found", http.StatusNotFound)
			return
		}
		cache.AddToCache(*order)
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(order); err != nil {
		logger.Error("Failed to encode order", zap.String("order_uid", orderUID), zap.Error(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	logger.Info("Served order", zap.String("order_uid", orderUID))
}
