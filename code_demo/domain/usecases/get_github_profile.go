package usecases

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"prometheus_and_Grafana_presentation/code_demo/configurations"
	"prometheus_and_Grafana_presentation/code_demo/domain/entities"
)

func GetGitHubProfile(username string) (entities.GitHubResponse, error) {
	github_url := configurations.ServiceConf.GithubUrl + username

	response, err := http.Get(github_url)

	var gitHubResposne entities.GitHubResponse

	if err != nil {
		return gitHubResposne, nil
	}

	data, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal(data, &gitHubResposne)
	return gitHubResposne, nil
}
