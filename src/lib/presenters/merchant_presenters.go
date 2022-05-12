package presenters

import (
	"strconv"

	"github.com/dreamingkills/steep/graph/model"
	"github.com/dreamingkills/steep/models"
)

func PresentMerchant(merchant models.Merchant) *model.Merchant {
	return &model.Merchant{
		ID:   strconv.Itoa(int(merchant.ID)),
		Name: merchant.Name,
		URL:  &merchant.URL,
	}
}
