package event

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"

	"github.com/samarec1812/audit-service/internal/app/entity"
)

const (
	historyTable = "history"
)

type Repository struct {
	db *sql.DB
}

func NewEventRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Save(ctx context.Context, event entity.Event) error {
	query, args, err := sq.Insert(historyTable).SetMap(event.GetEventDBRecord()).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
