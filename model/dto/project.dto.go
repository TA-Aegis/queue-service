package dto

type Project struct {
	ID       string `json:"id"`
	TenantID string `json:"tenant_id"`
	Name     string `json:"name"`
}

type CreateProjectRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CreateProjectResponse struct {
	Project
}
