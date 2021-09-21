package usecases

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"prometheus_and_Grafana_presentation/code_demo/configurations"
	"prometheus_and_Grafana_presentation/code_demo/domain/entities"
)

func GetGitLabProfile(username string) (entities.GitlabResponse, error) {
	gitlab_url := configurations.ServiceConf.GitlabUrl + username

	response, err := http.Get(gitlab_url)

	var gitLabResposne entities.GitlabResponse

	if err != nil {
		return gitLabResposne, nil
	}

	data, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal(data, &gitLabResposne)
	return gitLabResposne, nil
}
