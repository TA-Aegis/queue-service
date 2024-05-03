package infra

import (
	"antrein/bc-dashboard/model/config"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Repository struct {
	cfg *config.Config
}

func New(cfg *config.Config) *Repository {
	return &Repository{
		cfg: cfg,
	}
}

type InfraBody struct {
	ProjectID     string `json:"project_id"`
	ProjectDomain string `json:"project_domain"`
	URLPath       string `json:"url_path"`
}

func (r *Repository) GetInfraProjects(client *http.Client) ([]string, error) {
	resp, err := client.Get(r.cfg.Infra.ManagerURL + "/kube/project")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Data []string `json:"data"`
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (r *Repository) CreateInfraProject(client *http.Client, infraBody InfraBody) error {
	jsonData, err := json.Marshal(infraBody)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", r.cfg.Infra.ManagerURL+"/kube/project", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to create project, status code: %d", resp.StatusCode)
	}
	return nil
}

func (r *Repository) DeleteInfraProject(client *http.Client, projectId string) error {
	req, err := http.NewRequest("DELETE", r.cfg.Infra.ManagerURL+"/kube/project/"+projectId, nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to delete project, status code: %d", resp.StatusCode)
	}
	return nil
}
