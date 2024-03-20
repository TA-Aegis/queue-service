package main

import (
	"antrein/bc-dashboard/model/config"
	"context"
	_ "database/sql"

	"github.com/doug-martin/goqu/v9"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
)

type CommonResource struct {
	Db       *sqlx.DB
	QBuilder *goqu.DialectWrapper
	Vld      *validator.Validate
}

func NewCommonResource(cfg *config.Config, ctx context.Context) (*CommonResource, error) {
	db, err := sqlx.Open("postgres", cfg.Database.PostgreDB.Host)
	if err != nil {
		return nil, err
	}
	dialect := goqu.Dialect("postgres")
	vld := validator.New()

	rsc := CommonResource{
		Db:       db,
		QBuilder: &dialect,
		Vld:      vld,
	}
	return &rsc, nil
}
