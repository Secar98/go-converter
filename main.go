package main

import (
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/secar98/go-converter/cmd"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	_, err := os.ReadDir("uploads")
	if err != nil {
		log.Println("Creating uploads directory")
		os.Mkdir("uploads", 0777)
	}

	http.HandleFunc("POST /convert", cmd.ConvertHandler)
	http.HandleFunc("POST /convert-img", cmd.ConvertImageHandler)
	http.Handle("GET /metrics", promhttp.Handler())
	prometheus.Register(cmd.HttpDuration)

	metricsRouter := cmd.PrometheusMiddleware(http.DefaultServeMux)

	log.Println("Starting server on :" + port)
	http.ListenAndServe(":"+port, metricsRouter)
}
