package teas

import (
	"context"

	"github.com/dreamingkills/steep/graph/model"
	"github.com/dreamingkills/steep/lib/presenters"
	"github.com/dreamingkills/steep/models"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
)

func GetTeaByID(db *pgxpool.Pool, input *model.TeaInput) (*model.Tea, error) {
	var tea *models.Tea

	if err := db.QueryRow(context.Background(),
		`SELECT
            tea.id,
            tea.name,
            tea.type,
            merchant.id AS "merchant.id",
			merchant.name AS "merchant.name",
			merchant.url AS "merchant.url"
        FROM
            tea
        LEFT JOIN merchant ON merchant.id = tea.merchant_id
        WHERE
            tea.id = $1`, input.ID).Scan(&tea.ID, &tea.Name, &tea.Type, &tea.Merchant.ID, &tea.Merchant.Name, &tea.Merchant.URL); err != nil {
		if pgxscan.NotFound((err)) {
			return nil, nil
		} else {
			return nil, err
		}
	}

	if tea == nil {
		return nil, nil
	}

	return presenters.PresentTea(*tea), nil
}
