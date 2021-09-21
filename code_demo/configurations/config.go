package configurations

type AppConfig struct {
	Port int `yaml:"port"`
}

type ServiceConfig struct {
	GithubUrl string `yaml:"github_url"`
	GitlabUrl string `yaml:"gitlab_url"`
}

type MetricsConfig struct {
	MetricsNamespace string `yaml:"metrics_namespace"`
	MetricsSubsystem string `yaml:"metrics_subsystem"`
	Port             int    `yaml:"port"`
}
