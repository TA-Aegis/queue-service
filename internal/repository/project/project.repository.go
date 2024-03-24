package project

import (
	"antrein/bc-dashboard/model/config"
	"antrein/bc-dashboard/model/entity"
	"context"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	cfg *config.Config
	db  *sqlx.DB
}

func New(cfg *config.Config, db *sqlx.DB) *Repository {
	return &Repository{
		cfg: cfg,
		db:  db,
	}
}

func (r *Repository) CreateNewProject(ctx context.Context, req entity.Project) (*entity.Project, error) {
	project := req
	q := `INSERT INTO projects (id, name, tenant_id, created_at) VALUES ($1, $2, $3, $4)`
	_, err := r.db.ExecContext(ctx, q, req.ID, req.Name, req.TenantID, req.CreatedAt)
	return &project, err
}

func (r *Repository) GetTenantByID(ctx context.Context, id string) (*entity.Project, error) {
	project := entity.Project{}
	q := `SELECT * FROM projects WHERE id = $1 LIMIT 1`
	err := r.db.GetContext(ctx, &project, q, id)
	if err != nil {
		return nil, err
	}
	return &project, err
}

func (r *Repository) GetTenantProjectByID(ctx context.Context, id, tenantID string) (*entity.Project, error) {
	project := entity.Project{}
	q := `SELECT * FROM projects WHERE id = $1 AND tenant_id = $2 LIMIT 1`
	err := r.db.GetContext(ctx, &project, q, id, tenantID)
	if err != nil {
		return nil, err
	}
	return &project, err
}

func (r *Repository) GetTenantProjects(ctx context.Context, tenantID string) ([]entity.Project, error) {
	projects := []entity.Project{}
	q := `SELECT * FROM projects WHERE tenant_id = $1 ORDER BY id`
	err := r.db.SelectContext(ctx, &projects, q, tenantID)
	return projects, err
}

func (r *Repository) GetProjects(ctx context.Context, page int, pageSize int) ([]entity.Project, error) {
	projects := []entity.Project{}
	q := `SELECT * FROM projects ORDER BY name LIMIT $1 OFFSET $2`
	offset := (page - 1) * pageSize
	err := r.db.SelectContext(ctx, &projects, q, pageSize, offset)
	return projects, err
}
