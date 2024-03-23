package dto

type NewTenant struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type CreateTenantRequest struct {
	Email          string `json:"email"`
	Name           string `json:"name"`
	Password       string `json:"password"`
	RetypePassword string `json:"retype_password"`
}

type CreateTenantResponse struct {
	Tenant NewTenant `json:"tenant"`
	Token  string    `json:"token"`
}
