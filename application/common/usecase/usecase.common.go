package usecase

import (
	"antrein/bc-dashboard/application/common/repository"
	"antrein/bc-dashboard/internal/usecase/auth"
	"antrein/bc-dashboard/internal/usecase/project"
	"antrein/bc-dashboard/model/config"
)

type CommonUsecase struct {
	AuthUsecase    *auth.Usecase
	ProjectUsecase *project.Usecase
}

func NewCommonUsecase(cfg *config.Config, repo *repository.CommonRepository) (*CommonUsecase, error) {
	authUsecase := auth.New(cfg, repo.TenantRepo)
	projectUsecase := project.New(cfg, repo.ProjectRepo)

	commonUC := CommonUsecase{
		AuthUsecase:    authUsecase,
		ProjectUsecase: projectUsecase,
	}
	return &commonUC, nil
}
