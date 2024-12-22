package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var pingCounter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "ping_requests_total",
		Help: "Total number of ping requests",
	},
	[]string{"status"},
)

func init() {
	prometheus.MustRegister(pingCounter)
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	pingCounter.WithLabelValues("success").Inc()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}

func main() {
	http.HandleFunc("/ping", pingHandler)
	http.Handle("/metrics", promhttp.Handler())

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
