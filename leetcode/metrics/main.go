package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var requestCounter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "api_requests_total",
		Help: "Total number of API requests by client and request type.",
	},
	[]string{"client", "request_type"},
)

func init() {
	// Регистрация метрики
	prometheus.MustRegister(requestCounter)
}

func recordMetrics(client string, requestType string) {
	requestCounter.With(prometheus.Labels{"client": client, "request_type": requestType}).Inc()
}

func main() {
	// Пример использования
	http.HandleFunc("/some-endpoint", func(w http.ResponseWriter, r *http.Request) {
		client := r.Header.Get("X-Client-Id") // Получение идентификатора клиента из заголовка
		requestType := r.Method               // Тип запроса, например, GET или POST
		recordMetrics(client, requestType)

		w.Write([]byte("Request received"))
	})

	// Обработчик для экспонирования метрик
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}
