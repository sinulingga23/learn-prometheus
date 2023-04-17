package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sinulingga23/learn-prometheus/api"
	"github.com/sinulingga23/learn-prometheus/monitoring"
)

func main() {
	api := api.NewAPI()

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Post("/api/v1/products", api.AddProduct)
	r.Get("/api/v1/products", api.GetProducts)
	r.Get("/api/v1/products/{id}", api.GetProduct)

	prometheus.MustRegister(monitoring.TotalRequestApi, monitoring.DurationRequestAPi)
	r.Handle("/metrics", promhttp.Handler())

	if errListenAndServe := http.ListenAndServe(":8085", r); errListenAndServe != nil {
		log.Fatalf("Err Listen and Serve: %v", errListenAndServe)
	}
}
