package graph

import (
	"context"

	"github.com/dreamingkills/steep/graph/generated"
	"github.com/dreamingkills/steep/graph/model"
	"github.com/dreamingkills/steep/lib/presenters"
	"github.com/dreamingkills/steep/models"
	"github.com/georgysavva/scany/pgxscan"
)

func (r *mutationResolver) CreateMerchant(ctx context.Context, input model.NewMerchant) (*model.Merchant, error) {
	var dbMerchant models.Merchant
	if err := r.DB.QueryRow(context.Background(),
		`INSERT INTO
			merchant (name, url) 
		VALUES ($1, $2) RETURNING *`, input.Name, input.URL).Scan(&dbMerchant.ID, &dbMerchant.Name, &dbMerchant.URL); err != nil {
		return nil, err
	}

	return presenters.PresentMerchant(dbMerchant), nil
}

func (r *mutationResolver) CreateTea(ctx context.Context, input model.CreateTeaInput) (*model.Tea, error) {
	var dbTea models.Tea

	if err := r.DB.QueryRow(context.Background(),
		`INSERT INTO 
			tea (name, type, merchant_id) 
		VALUES ($1, $2, $3) RETURNING *`, input.Name, input.Type.String(), input.MerchantID).Scan(&dbTea.ID, &dbTea.Name, &dbTea.Type, &dbTea.MerchantID); err != nil {
		return nil, err
	}

	if err := r.DB.QueryRow(context.Background(), "SELECT id, name, url FROM merchant WHERE id = $1", dbTea.MerchantID).Scan(&dbTea.Merchant.ID, &dbTea.Merchant.Name, &dbTea.Merchant.URL); err != nil {
		return nil, err
	}

	return presenters.PresentTea(dbTea), nil
}

func (r *queryResolver) Merchant(ctx context.Context, input *model.MerchantInput) (*model.Merchant, error) {
	if input.ID == nil && input.Name == nil {
		return nil, nil
	}

	var dbMerchant models.Merchant

	if err := r.DB.QueryRow(context.Background(),
		`SELECT 
			id,
			name,
			url
		FROM
			merchant
		WHERE
			name = $1 
			OR id = $2`, input.Name, input.ID).Scan(&dbMerchant.ID, &dbMerchant.Name, &dbMerchant.URL); err != nil {
		if pgxscan.NotFound(err) {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return presenters.PresentMerchant(dbMerchant), nil
}

func (r *queryResolver) Tea(ctx context.Context, input *model.TeaInput) (*model.Tea, error) {
	var dbTea models.Tea

	if err := r.DB.QueryRow(context.Background(),
		`SELECT 
			tea.id,
			tea.name,
			tea.type,
			merchant.id AS "merchant.id",
			merchant.name AS "merchant.name",
			merchant.url AS "merchant.url"
		FROM tea 
		LEFT JOIN merchant ON merchant.id = tea.merchant_id
		WHERE tea.id = $1`, input.ID).Scan(&dbTea.ID, &dbTea.Name, &dbTea.Type, &dbTea.Merchant.ID, &dbTea.Merchant.Name, &dbTea.Merchant.URL); err != nil {
		if pgxscan.NotFound(err) {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return presenters.PresentTea(dbTea), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
