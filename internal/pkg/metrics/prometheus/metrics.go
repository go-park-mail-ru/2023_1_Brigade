package prometheus

import (
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
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
	e := echo.New()
	err := e.Start(address)
	return err
}
