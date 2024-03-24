package configuration

import (
	"antrein/bc-dashboard/internal/repository/configuration"
	"antrein/bc-dashboard/model/config"
	"antrein/bc-dashboard/model/dto"
	"antrein/bc-dashboard/model/entity"
	"antrein/bc-dashboard/model/types"
	"context"
	"database/sql"
	"log"
)

type Usecase struct {
	cfg  *config.Config
	repo *configuration.Repository
}

func New(cfg *config.Config, repo *configuration.Repository) *Usecase {
	return &Usecase{
		cfg:  cfg,
		repo: repo,
	}
}

func (u *Usecase) UpdateProjectConfig(ctx context.Context, req dto.UpdateProjectConfig) *dto.ErrorResponse {
	var errRes dto.ErrorResponse

	config := entity.Configuration{
		ProjectID:   req.ProjectID,
		Threshold:   req.Threshold,
		SessionTime: req.SessionTime,
		Host: sql.NullString{
			Valid:  true,
			String: req.Host,
		},
		BaseURL: sql.NullString{
			Valid:  true,
			String: req.BaseURL,
		},
		MaxUsersInQueue: req.MaxUsersInQueue,
		PagesToApply: types.NullStringArray{
			Valid:       true,
			StringArray: req.PagesToApply,
		},
		QueueStart: sql.NullTime{
			Valid: true,
			Time:  req.QueueStart,
		},
		QueueEnd: sql.NullTime{
			Valid: true,
			Time:  req.QueueEnd,
		},
	}

	err := u.repo.UpdateProjectConfig(ctx, config)
	if err != nil {
		log.Println("Error gagal mengupdate konfigurasi project", err)
		if err == sql.ErrNoRows {
			errRes = dto.ErrorResponse{
				Status: 404,
				Error:  "Project dengan id tersebut tidak ditemukan",
			}
			return &errRes
		}
		errRes = dto.ErrorResponse{
			Status: 500,
			Error:  "Gagal mengupdate konfigurasi project",
		}
		return &errRes
	}

	return nil
}

func (u *Usecase) UpdateProjectStyle(ctx context.Context, req dto.UpdateProjectStyle) *dto.ErrorResponse {
	var errRes dto.ErrorResponse

	config := entity.Configuration{
		ProjectID:      req.ProjectID,
		QueuePageStyle: req.QueuePageStyle,
		QueueHTMLPage: sql.NullString{
			Valid:  true,
			String: req.QueueHTMLPage,
		},
		QueuePageBaseColor: sql.NullString{
			Valid:  true,
			String: req.QueuePageBaseColor,
		},
		QueuePageTitle: sql.NullString{
			Valid:  true,
			String: req.QueuePageTitle,
		},
		QueuePageLogo: sql.NullString{
			Valid:  true,
			String: req.QueuePageLogo,
		},
	}

	err := u.repo.UpdateProjectConfig(ctx, config)
	if err != nil {
		log.Println("Error gagal mengupdate tampilan project", err)
		if err == sql.ErrNoRows {
			errRes = dto.ErrorResponse{
				Status: 404,
				Error:  "Project dengan id tersebut tidak ditemukan",
			}
			return &errRes
		}
		errRes = dto.ErrorResponse{
			Status: 500,
			Error:  "Gagal mengupdate tampilan project",
		}
		return &errRes
	}

	return nil
}
