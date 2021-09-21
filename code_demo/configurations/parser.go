package configurations

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

var Appconf AppConfig
var ServiceConf ServiceConfig
var MetricsConf MetricsConfig

func Load() {
	Appconf = parseAppConfig()
	ServiceConf = parseServiceConfig()
	MetricsConf = parseMetricsConfig()
}

// parse configurations in app.yaml
func parseAppConfig() AppConfig {

	content := read("app.yaml")

	cfg := AppConfig{}

	err := yaml.Unmarshal(content, &cfg)

	if err != nil {
		panic(fmt.Sprintf("error: %v", err))
	}

	return cfg
}

// parse configurations in services.yaml
func parseServiceConfig() ServiceConfig {

	content := read("service.yaml")

	cfg := ServiceConfig{}

	err := yaml.Unmarshal(content, &cfg)

	if err != nil {
		panic(fmt.Sprintf("error: %v", err))
	}

	return cfg
}

// parse configurations in metrics.yaml
func parseMetricsConfig() MetricsConfig {

	content := read("metrics.yaml")

	cfg := MetricsConfig{}

	err := yaml.Unmarshal(content, &cfg)

	if err != nil {
		panic(fmt.Sprintf("error: %v", err))
	}

	return cfg
}
