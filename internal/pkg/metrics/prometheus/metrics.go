package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type MetricsGRPC struct {
	MicroserviceName string
	Hits             prometheus.Counter
	ResponseTime     *prometheus.HistogramVec
}

func NewMetricsGRPCServer(microserviceName string) (*MetricsGRPC, error) {
	metrics := &MetricsGRPC{
		Hits: prometheus.NewCounter(prometheus.CounterOpts{
			Name: microserviceName + "_hits",
			Help: "счетчик запрос на микросервис",
		}),
		ResponseTime: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Name: microserviceName + "_timings",
			Help: "время ответа запроса от микросервиса",
		}, []string{"code", "method"}),
		MicroserviceName: microserviceName,
	}

	if err := prometheus.Register(metrics.Hits); err != nil {
		return nil, err
	}

	if err := prometheus.Register(metrics.ResponseTime); err != nil {
		return nil, err
	}

	return metrics, nil
}

func (m MetricsGRPC) StartGRPCMetricsServer(address string) error {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	log.Info(m.MicroserviceName, ": starting metrics server...")

	return http.ListenAndServe(address, mux)
}
