package presenters

import (
	"strconv"

	"github.com/dreamingkills/steep/graph/model"
	"github.com/dreamingkills/steep/models"
)

func PresentTea(tea models.Tea) *model.Tea {
	return &model.Tea{
		ID:       strconv.Itoa(int(tea.ID)),
		Name:     tea.Name,
		Type:     model.TeaType(tea.Type),
		Merchant: PresentMerchant(tea.Merchant),
	}
}
