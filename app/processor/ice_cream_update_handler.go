package processor

import (
	"context"
	"ic-indexer-service/app/helpers"
	"ic-indexer-service/app/model/bo"
	"ic-indexer-service/app/model/request"
	"log"
)

//go:generate sh -c "$GOPATH/bin/mockery -case=underscore -dir=. -name=UpdateHandler"
type UpdateHandler interface {
	UpdateIcecreamDetails(ctx context.Context, icecream *bo.ESIcecream, params request.IcecreamIndexRequest) error
}

type updateHandler struct {
}

func NewUpdateHandler() UpdateHandler {
	return &updateHandler{}
}

func (updateHandler) UpdateIcecreamDetails(ctx context.Context, icecream *bo.ESIcecream, params request.IcecreamIndexRequest) error {

	log.Print(ctx, "updating icecream details")

	icecream.Id = helpers.GetInt64(params.Id, icecream.Id)
	icecream.Name = helpers.GetString(params.Name, icecream.Name)
	icecream.ProductId = helpers.GetString(params.ProductId, icecream.ProductId)
	icecream.Ingredients = helpers.GetArrayString(params.Ingredients, icecream.Ingredients)
	icecream.SourcingValues = helpers.GetArrayString(params.SourcingValues, icecream.SourcingValues)
	icecream.DietaryCertifications = helpers.GetString(params.DietaryCertifications, icecream.DietaryCertifications)
	icecream.Description = helpers.GetString(params.Description, icecream.Description)
	icecream.Story = helpers.GetString(params.Story, icecream.Story)
	icecream.ImageOpen = helpers.GetString(params.ImageOpen, icecream.ImageOpen)
	icecream.ImageClosed = helpers.GetString(params.ImageClosed, icecream.ImageClosed)
	icecream.AllergyInfo = helpers.GetString(params.AllergyInfo, icecream.AllergyInfo)

	return nil

}
