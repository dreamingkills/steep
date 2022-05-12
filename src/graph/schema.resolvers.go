package graph

import (
	"context"

	"github.com/dreamingkills/steep/graph/generated"
	"github.com/dreamingkills/steep/graph/model"
	"github.com/dreamingkills/steep/lib/presenters"
	"github.com/dreamingkills/steep/lib/services/merchants"
	"github.com/dreamingkills/steep/lib/services/teas"
	"github.com/dreamingkills/steep/models"
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

	dbMerchant, err := merchants.GetMerchantByID(r.DB, dbTea.MerchantID)
	if err != nil {
		return nil, err
	}

	dbTea.Merchant = *dbMerchant

	return presenters.PresentTea(dbTea), nil
}

func (r *queryResolver) Merchant(ctx context.Context, input *model.MerchantInput) (*model.Merchant, error) {
	merchant, err := merchants.GetMerchant(r.DB, input)

	if err != nil {
		return nil, err
	}

	return merchant, nil
}

func (r *queryResolver) Tea(ctx context.Context, input *model.TeaInput) (*model.Tea, error) {
	tea, err := teas.GetTeaByID(r.DB, input)

	if err != nil {
		return nil, err
	}

	return tea, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
