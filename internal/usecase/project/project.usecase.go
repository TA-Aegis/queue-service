package project

import (
	"antrein/bc-dashboard/internal/repository/project"
	"antrein/bc-dashboard/model/config"
	"antrein/bc-dashboard/model/dto"
	"antrein/bc-dashboard/model/entity"
	"context"
	"log"
	"time"

	"github.com/lib/pq"
)

type Usecase struct {
	cfg  *config.Config
	repo *project.Repository
}

func New(cfg *config.Config, repo *project.Repository) *Usecase {
	return &Usecase{
		cfg:  cfg,
		repo: repo,
	}
}

func (u *Usecase) RegisterNewProject(ctx context.Context, req dto.CreateProjectRequest, tenantID string) (*dto.CreateProjectResponse, *dto.ErrorResponse) {
	var errRes dto.ErrorResponse

	project := entity.Project{
		ID:        req.ID,
		Name:      req.Name,
		TenantID:  tenantID,
		CreatedAt: time.Now(),
	}

	created, err := u.repo.CreateNewProject(ctx, project)
	if err != nil {
		log.Println("Error gagal membuat project", err)
		if pgErr, ok := err.(*pq.Error); ok && pgErr.Code == "23505" {
			errRes = dto.ErrorResponse{
				Status: 400,
				Error:  "Project dengan id tersebut sudah ada",
			}
			return nil, &errRes
		}
		errRes = dto.ErrorResponse{
			Status: 500,
			Error:  "Gagal membuat project",
		}
		return nil, &errRes
	}

	return &dto.CreateProjectResponse{
		Project: dto.Project{
			ID:       created.ID,
			Name:     created.Name,
			TenantID: created.TenantID,
		},
	}, nil
}
