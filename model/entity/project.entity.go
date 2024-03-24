package entity

import (
	"database/sql"
	"time"
)

type Project struct {
	ID        string       `db:"id"`
	Name      string       `db:"name"`
	TenantID  string       `db:"tenant_id"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at,omitempty"`
}
