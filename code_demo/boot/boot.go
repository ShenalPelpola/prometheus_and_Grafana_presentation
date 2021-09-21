package boot

import (
	"prometheus_and_Grafana_presentation/code_demo/configurations"
	"prometheus_and_Grafana_presentation/code_demo/http"
	"prometheus_and_Grafana_presentation/code_demo/metrics"
)

func Init() {
	configurations.Load()

	go metrics.Init()

	http.InitServer()
}
