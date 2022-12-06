package builder

import (
	"ic-indexer-service/app/model/bo"
	"ic-indexer-service/app/model/response"
)

func IcecreamIndexResponseBuilder(icecream bo.ESIcecream) response.IcecreamIndexResponse {
	var icecreamResponse response.IcecreamIndexResponse

	icecreamResponse.ProductId = icecream.ProductId
	icecreamResponse.Name = icecream.Name
	icecreamResponse.AllergyInfo = icecream.AllergyInfo
	icecreamResponse.Description = icecream.Description
	icecreamResponse.Id = icecream.Id
	icecreamResponse.ImageClosed = icecream.ImageClosed
	icecreamResponse.ImageOpen = icecream.ImageOpen
	icecreamResponse.Ingredients = icecream.Ingredients
	icecreamResponse.SourcingValues = icecream.SourcingValues
	icecreamResponse.Story = icecream.Story
	icecreamResponse.DietaryCertifications = icecream.DietaryCertifications

	return icecreamResponse
}
