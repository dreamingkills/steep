package merchants

import (
	"context"
	"strconv"

	"github.com/dreamingkills/steep/graph/model"
	"github.com/dreamingkills/steep/lib/presenters"
	"github.com/dreamingkills/steep/models"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
)

func GetMerchantByID(db *pgxpool.Pool, id uint) (*models.Merchant, error) {
	var merchant models.Merchant

	if err := db.QueryRow(context.Background(),
		`SELECT
            id,
            name,
            url
        FROM
            merchant
        WHERE
            id = $1`, id).Scan(&merchant.ID, &merchant.Name, &merchant.URL); err != nil {
		if pgxscan.NotFound((err)) {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return &merchant, nil
}

func GetMerchantByName(db *pgxpool.Pool, name string) (*models.Merchant, error) {
	var merchant models.Merchant

	if err := db.QueryRow(context.Background(),
		`SELECT
			id,
			name,
			url
		FROM
			merchant
		WHERE
			name = $1`, name).Scan(&merchant.ID, &merchant.Name, &merchant.URL); err != nil {
		if pgxscan.NotFound((err)) {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return &merchant, nil
}

func GetMerchant(db *pgxpool.Pool, input *model.MerchantInput) (*model.Merchant, error) {
	var merchant *models.Merchant

	if input.ID != nil {
		id, err := strconv.ParseUint(*input.ID, 10, 0)
		if err != nil {
			return nil, err
		}

		dbMerchant, err := GetMerchantByID(db, uint(id))
		if err != nil {
			return nil, err
		} else if dbMerchant == nil {
			return nil, nil
		}

		merchant = dbMerchant
	} else if input.Name != nil {
		dbMerchant, err := GetMerchantByName(db, *input.Name)
		if err != nil {
			return nil, err
		} else if dbMerchant == nil {
			return nil, nil
		}

		merchant = dbMerchant
	} else {
		return nil, nil
	}

	if merchant == nil {
		return nil, nil
	}

	return presenters.PresentMerchant(*merchant), nil
}
