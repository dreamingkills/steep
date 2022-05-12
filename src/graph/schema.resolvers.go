package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"strconv"

	"github.com/dreamingkills/steep/graph/generated"
	"github.com/dreamingkills/steep/graph/model"
	"github.com/dreamingkills/steep/models"
	"gorm.io/gorm"
)

func (r *mutationResolver) CreateMerchant(ctx context.Context, input model.NewMerchant) (*model.Merchant, error) {
	var url string

	if(input.URL != nil) {
		url = *input.URL
	}

	dbMerchant := models.Merchant{Name: input.Name, URL: url}
	r.DB.Save(&dbMerchant)

	merchant := model.Merchant{
		ID:   strconv.Itoa(int(dbMerchant.ID)),
		Name: dbMerchant.Name,
		URL:  &dbMerchant.URL,
	}

	return &merchant, nil
}

func (r *queryResolver) Merchant(ctx context.Context, input *model.MerchantInput) (*model.Merchant, error) {
	if input.ID == nil && input.Name == nil {
		return nil, nil
	}

	var id int

	if input.ID != nil {
		_id, err := strconv.Atoi(*input.ID)

		if err != nil {
			return nil, err
		}

		id = _id
	}

	var dbMerchant models.Merchant

	if err := r.DB.Where("name = ? OR id = ?", input.Name, id).First(&dbMerchant).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}

	merchant := model.Merchant{
		ID:   strconv.Itoa(int(dbMerchant.ID)),
		Name: dbMerchant.Name,
		URL:  &dbMerchant.URL,
	}

	return &merchant, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
