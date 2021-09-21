package http

import (
	"fmt"
	"net/http"
	"prometheus_and_Grafana_presentation/code_demo/configurations"
	"prometheus_and_Grafana_presentation/code_demo/http/handlers"

	"github.com/gorilla/mux"
)

func InitServer() {
	r := mux.NewRouter()

	// r.HandleFunc("/api/gitHubProfile", handlers.GetGithubProfile).Methods("GET", "POST")
	r.HandleFunc("/api/gitLabProfile", handlers.GetGitlabProfile).Methods("GET", "POST")

	err := http.ListenAndServe(fmt.Sprintf(`:%d`, configurations.Appconf.Port), r)
	if err != nil {
		fmt.Printf(`%s`, err)
	}
}
