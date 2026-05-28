package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func InsertRow(ctx context.Context, conn *pgx.Conn) error {
	sqlQuery := `
	INSERT INTO tasks (title, description, completed, created_at)
	VALUES ('Выгул собаки', 'Сделать ДЗ по матеше', FALSE, '2026-05-28 14:00:00');
	`
	_, err := conn.Exec(ctx, sqlQuery)

	return err
}
