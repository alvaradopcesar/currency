package repository

import (
	"context"
	"database/sql"
	"errors"
	"strconv"
	"strings"
	"time"
)

type Repository interface {
	GetAll(ctx context.Context) ([]Currency, error)
	GetByID(ctx context.Context, f CurrencyFilter) ([]Currency, error)
	Insert(ctx context.Context, req *InsertCurrency) error
	InsertQuery(ctx context.Context, req *InsertQuery) error
}

type postgres struct {
	db      *sql.DB
	timeOut time.Duration
}

func NewRepository(db *sql.DB, timeOut time.Duration) Repository {
	return &postgres{
		db:      db,
		timeOut: timeOut * time.Second,
	}
}

func (p postgres) GetAll(ctx context.Context) ([]Currency, error) {
	q := `select DISTINCT ON (code) code,"value", created_at from currency order by code, created_at desc;`

	ctx, cancel := context.WithTimeout(ctx, p.timeOut)
	defer cancel()

	res := make([]Currency, 0)
	rows, err := p.db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var c Currency
		err = rows.Scan(
			&c.Code,
			&c.Value,
			&c.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		res = append(res, c)
	}

	return res, nil
}

func (p postgres) GetByID(ctx context.Context, f CurrencyFilter) ([]Currency, error) {
	var q strings.Builder

	q.WriteString(`SELECT customer_id, code, "value", created_at FROM currency`)
	args := make([]interface{}, 0)
	q.WriteString(" WHERE")
	if f.Code != "" {
		args = append(args, f.Code)
		q.WriteString(" code = $")
		q.WriteString(strconv.Itoa(len(args)))
	}

	if !f.FEnd.IsZero() {
		if len(args) > 0 {
			q.WriteString(" AND")
		}
		args = append(args, f.FInit, f.FEnd)
		q.WriteString(" created_at BETWEEN $")
		q.WriteString(strconv.Itoa(len(args) - 1))
		q.WriteString(` AND $`)
		q.WriteString(strconv.Itoa(len(args)))
	}

	ctx, cancel := context.WithTimeout(ctx, p.timeOut)
	defer cancel()

	res := make([]Currency, 0)
	rows, err := p.db.QueryContext(ctx, q.String(), args...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var c Currency
		err = rows.Scan(
			&c.CustomerID,
			&c.Code,
			&c.Value,
			&c.CreatedAt,
		)
		if err != nil {
			continue
		}

		res = append(res, c)
	}

	return res, nil
}

func (p postgres) Insert(ctx context.Context, req *InsertCurrency) error {
	q := `INSERT INTO currency (code, "value") VALUES($1, $2);`

	ctx, cancel := context.WithTimeout(ctx, p.timeOut)
	defer cancel()

	stmt, err := p.db.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.ExecContext(ctx,
		req.Code,
		req.Value,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected != 1 {
		return errors.New("expected affects a row")
	}

	return nil
}

func (p postgres) InsertQuery(ctx context.Context, req *InsertQuery) error {
	q := `INSERT INTO query ("method", address, code, "time") VALUES($1, $2, $3, $4);`

	ctx, cancel := context.WithTimeout(ctx, p.timeOut)
	defer cancel()

	stmt, err := p.db.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.ExecContext(ctx,
		req.Method,
		req.Address,
		req.Code,
		req.Time,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected != 1 {
		return errors.New("expected affects a row")
	}

	return nil
}
