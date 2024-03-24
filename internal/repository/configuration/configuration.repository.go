package configuration

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

func (r *Repository) UpdateProjectConfig(ctx context.Context, req entity.Configuration) error {
	q := `UPDATE configurations 
		  SET threshold = $1, 
		  session_time = $2, 
		  host = $3, 
		  base_url = $4,
		  max_users_in_queue = $5,
		  pages_to_apply = $6,
		  queue_start = $7,
		  queue_end = $8,
		  is_configure = $9,
		  updated_at = now()
		  WHERE project_id = $10`
	_, err := r.db.ExecContext(ctx, q, req.Threshold, req.SessionTime, req.Host, req.BaseURL, req.MaxUsersInQueue, req.PagesToApply, req.QueueStart, req.QueueEnd, true, req.ProjectID)
	return err
}

func (r *Repository) UpdateProjectStyle(ctx context.Context, req entity.Configuration) error {
	q := `UPDATE configurations 
		  SET queue_page_style = $1,
		  queue_html_page = $2,
		  queue_page_base_color = $3,
		  queue_page_title = $4,
		  queue_page_logo = $5,
		  updated_at = now()
		  WHERE project_id = $6`
	_, err := r.db.ExecContext(ctx, q, req.QueuePageStyle, req.QueueHTMLPage, req.QueuePageBaseColor, req.QueuePageTitle, req.QueuePageLogo, req.ProjectID)
	return err
}
