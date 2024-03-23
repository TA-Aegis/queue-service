package repository

import (
	"antrein/bc-dashboard/application/common/resource"
	"antrein/bc-dashboard/internal/repository/tenant"
	"antrein/bc-dashboard/model/config"
)

type CommonRepository struct {
	TenantRepo *tenant.Repository
}

func NewCommonRepository(cfg *config.Config, rsc *resource.CommonResource) (*CommonRepository, error) {
	tenantRepo := tenant.New(cfg, rsc.Db)

	commonRepo := CommonRepository{
		TenantRepo: tenantRepo,
	}
	return &commonRepo, nil
}
