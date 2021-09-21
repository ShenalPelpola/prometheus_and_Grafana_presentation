package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"prometheus_and_Grafana_presentation/code_demo/domain/usecases"
	"prometheus_and_Grafana_presentation/code_demo/metrics"
	"time"

	"github.com/pickme-go/log"
)

type Response struct {
	Value interface{} `json:"data"`
}

type User struct {
	Username string `json:"username"`
}

func GetGithubProfile(w http.ResponseWriter, r *http.Request) error {
	user := User{}
	var username string
	start := time.Now()

	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)

		err := decoder.Decode(&user)
		if err != nil {
			return err
		}
		username = user.Username
	} else if r.Method == "GET" {
		username = r.FormValue("username")
	}

	response, err := usecases.GetGitHubProfile(username)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}

	res := Response{
		Value: response,
	}

	metrics.IncrementRequestCount(r.Method, "/api/gitHubProfile")
	metrics.SetTotalLatency(start, r.Method, "/api/gitHubProfile")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(res)
}

func GetGitlabProfile(w http.ResponseWriter, r *http.Request) {
	user := User{}
	var username string
	start := time.Now()

	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		log.Info("body of the request : ", decoder)

		err := decoder.Decode(&user)
		if err != nil {
			log.Error("ERROR decoding json: ", err)
			return
		}
		username = user.Username
	} else if r.Method == "GET" {
		username = r.FormValue("username")
	}

	response, err := usecases.GetGitLabProfile(username)
	if err != nil {
		log.Fatal("The HTTP request failed with error %s\n", err)
	}

	res := Response{
		Value: response,
	}

	metrics.SetTotalLatency(start, r.Method, "/api/gitLabProfile")
	metrics.IncrementRequestCount(r.Method, "/api/gitLabProfile")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
