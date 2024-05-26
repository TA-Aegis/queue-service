package configuration

import (
	"antrein/bc-dashboard/internal/repository/configuration"
	"antrein/bc-dashboard/internal/repository/infra"
	"antrein/bc-dashboard/model/config"
	"antrein/bc-dashboard/model/dto"
	"antrein/bc-dashboard/model/entity"
	"context"
	"database/sql"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"time"
)

type Usecase struct {
	cfg       *config.Config
	repo      *configuration.Repository
	infraRepo *infra.Repository
}

func New(cfg *config.Config, repo *configuration.Repository, infraRepo *infra.Repository) *Usecase {
	return &Usecase{
		cfg:       cfg,
		repo:      repo,
		infraRepo: infraRepo,
	}
}

func (u *Usecase) GetProjectConfigByID(ctx context.Context, projectID string) (*dto.ProjectConfig, *dto.ErrorResponse) {
	var errRes dto.ErrorResponse

	config, err := u.repo.GetConfigByProjectID(ctx, projectID)
	if err != nil {
		if err == sql.ErrNoRows {
			errRes = dto.ErrorResponse{
				Status: 404,
				Error:  "Project dengan id tersebut tidak ditemukan",
			}
			return nil, &errRes
		}
		log.Println(err)
		errRes = dto.ErrorResponse{
			Status: 500,
			Error:  "Gagal mendapatkan konfigurasi project",
		}
		return nil, &errRes
	}
	return &dto.ProjectConfig{
		ProjectID:          config.ProjectID,
		Threshold:          config.Threshold,
		SessionTime:        config.SessionTime,
		Host:               config.Host.String,
		BaseURL:            config.BaseURL.String,
		MaxUsersInQueue:    config.MaxUsersInQueue,
		QueueStart:         config.QueueStart.Time,
		QueueEnd:           config.QueueEnd.Time,
		QueuePageStyle:     config.QueuePageStyle,
		QueueHTMLPage:      config.QueueHTMLPage.String,
		QueuePageBaseColor: config.QueuePageBaseColor.String,
		QueuePageTitle:     config.QueuePageTitle.String,
		QueuePageLogo:      config.QueuePageLogo.String,
	}, nil
}

func (u *Usecase) GetProjectConfigByHost(ctx context.Context, host string) (*dto.ProjectConfig, *dto.ErrorResponse) {
	var errRes dto.ErrorResponse

	config, err := u.repo.GetConfigByHost(ctx, host)
	if err != nil {
		if err == sql.ErrNoRows {
			errRes = dto.ErrorResponse{
				Status: 404,
				Error:  "Project dengan host tersebut tidak ditemukan",
			}
			return nil, &errRes
		}
		log.Println(err)
		errRes = dto.ErrorResponse{
			Status: 500,
			Error:  "Gagal mendapatkan konfigurasi project",
		}
		return nil, &errRes
	}
	return &dto.ProjectConfig{
		ProjectID:          config.ProjectID,
		Threshold:          config.Threshold,
		SessionTime:        config.SessionTime,
		Host:               config.Host.String,
		BaseURL:            config.BaseURL.String,
		MaxUsersInQueue:    config.MaxUsersInQueue,
		QueueStart:         config.QueueStart.Time,
		QueueEnd:           config.QueueEnd.Time,
		QueuePageStyle:     config.QueuePageStyle,
		QueueHTMLPage:      config.QueueHTMLPage.String,
		QueuePageBaseColor: config.QueuePageBaseColor.String,
		QueuePageTitle:     config.QueuePageTitle.String,
		QueuePageLogo:      config.QueuePageLogo.String,
	}, nil
}

func (u *Usecase) UpdateProjectConfig(ctx context.Context, req dto.UpdateProjectConfig) *dto.ErrorResponse {
	var errRes dto.ErrorResponse

	const layout = "2006-01-02T15:04:05"
	queueStart, err := time.Parse(layout, req.QueueStart)
	if err != nil {
		errRes = dto.ErrorResponse{
			Status: 400,
			Error:  "Format waktu queue mulai salah",
		}
		return &errRes
	}

	queueEnd, err := time.Parse(layout, req.QueueEnd)
	if err != nil {
		errRes = dto.ErrorResponse{
			Status: 400,
			Error:  "Format waktu queue berakhir salah",
		}
		return &errRes
	}

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
		QueueStart: sql.NullTime{
			Valid: true,
			Time:  queueStart,
		},
		QueueEnd: sql.NullTime{
			Valid: true,
			Time:  queueEnd,
		},
	}

	err = u.repo.UpdateProjectConfig(ctx, config)
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

func (u *Usecase) UpdateProjectStyle(ctx context.Context, req dto.UpdateProjectStyle, imageFile *multipart.FileHeader, htmlFile *multipart.FileHeader) *dto.ErrorResponse {
	var errRes dto.ErrorResponse

	if req.QueuePageStyle == "base" {
		if imageFile != nil {
			image, err := imageFile.Open()
			if err != nil {
			}
			imageContent, err := io.ReadAll(image)
			if err != nil {
			}
			err = u.infraRepo.UploadLogoFile(&http.Client{}, dto.File{
				Filename: req.ProjectID,
				Content:  imageContent,
			})
		}
	} else if req.QueuePageStyle == "custom" {

	} else {
		errRes = dto.ErrorResponse{
			Status: 400,
			Error:  "Tipe style tidak valid",
		}
		return &errRes
	}

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
