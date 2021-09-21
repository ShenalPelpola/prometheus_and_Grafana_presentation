package metrics

import (
	"fmt"
	"net/http"
	"prometheus_and_Grafana_presentation/code_demo/configurations"
	"time"

	"github.com/gorilla/mux"
	"github.com/pickme-go/log"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

/*
Steps of monitoring your applications with prometheus

1. Declaring the metrics, their type, and the labels
2. Initialize the metrics and give them labels(if you are using any)
3. Register each metric using prometheus.Register() function.
4. Initialize a separate router which runs on a different port from the application port
*/

// Declaring the metrics
var (
	requestCount   *prometheus.CounterVec
	requestLatency *prometheus.SummaryVec
	// labels are a way to filter out metrics
	labels = []string{`method`, `endpoint`}
)

func Init() {
	ok := make(chan bool)

	registerMetrics(configurations.MetricsConf.MetricsNamespace, configurations.MetricsConf.MetricsSubsystem)

	r := mux.NewRouter()
	r.Handle(`/metrics`, promhttp.Handler()).Methods(http.MethodGet)

	go func() {
		err := http.ListenAndServe(fmt.Sprintf(`:%d`, configurations.MetricsConf.Port), r)
		if err != nil {
			log.Fatal(`cannot start web server : `, err)
		}
	}()

	log.Info(fmt.Sprintf(`http server started on port %d`, configurations.MetricsConf.Port))
	<-ok
}

// This function initialize the defined metrics
func initMetrics(namespace string, subsystem string) {
	requestCount = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      `request_count`,
		Help:      `request_count`,
	}, labels)

	requestLatency = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      `request_latency`,
		Help:      `request latency`,
	}, labels)
}

func registerMetrics(namespace, subsystem string) {
	initMetrics(namespace, subsystem)

	err := prometheus.Register(requestCount)
	if err != nil {
		log.Fatal(`request count metrics registration failed`)
	}

	err = prometheus.Register(requestLatency)
	if err != nil {
		log.Fatal(`request latency metrics registration failed`)
	}
}

// Functions for setting metrics
// Thess methods are used throughout the application when a metrics is neeed to be updated
func IncrementRequestCount(method string, endpoint string) {
	labels := prometheus.Labels{"method": method, "endpoint": endpoint}
	requestCount.With(labels).Add(1)
}

func SetTotalLatency(start time.Time, method string, endpoint string) {
	labels := prometheus.Labels{"method": method, "endpoint": endpoint}
	requestLatency.With(labels).Observe(float64(time.Since(start).Nanoseconds() / 1000))
}
