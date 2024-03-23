package usecase

import (
	"antrein/bc-dashboard/application/common/repository"
	"antrein/bc-dashboard/internal/usecase/auth"
	"antrein/bc-dashboard/model/config"
)

type CommonUsecase struct {
	AuthUsecase *auth.Usecase
}

func NewCommonUsecase(cfg *config.Config, repo *repository.CommonRepository) (*CommonUsecase, error) {
	authUsecase := auth.New(cfg, repo.TenantRepo)

	commonUC := CommonUsecase{
		AuthUsecase: authUsecase,
	}
	return &commonUC, nil
}
