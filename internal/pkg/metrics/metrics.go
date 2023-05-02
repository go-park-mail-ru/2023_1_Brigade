package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type MetricsGRPC struct {
	Errors  *prometheus.CounterVec
	Hits    prometheus.Counter
	Timings *prometheus.HistogramVec
	Name    string
}

func NewMetricsGRPCServer(name string) (*MetricsGRPC, error) {
	metrics := &MetricsGRPC{
		Hits: prometheus.NewCounter(prometheus.CounterOpts{
			Name: name + "_hits",
			Help: "counts all hits for microservice",
		}),
		Errors: prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: name + "_errors",
			Help: "counts responses with error from microservice",
		}, []string{"code", "method"}),
		Timings: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Name: name + "_timings",
			Help: "measures duration of microservice method",
		}, []string{"code", "method"}),
		Name: name,
	}

	if err := prometheus.Register(metrics.Hits); err != nil {
		return nil, err
	}

	if err := prometheus.Register(metrics.Errors); err != nil {
		return nil, err
	}

	if err := prometheus.Register(metrics.Timings); err != nil {
		return nil, err
	}

	return metrics, nil
}

func (m MetricsGRPC) RunGRPCMetricsServer(address string) error {
	//use separated ServeMux to prevent handling on the global Mux
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	log.Info(m.Name, ": starting metrics server...")

	return http.ListenAndServe(address, mux)
}
