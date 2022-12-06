package response

type BulkIcecreamIndexResponse struct {
	Icecreams    []IcecreamIndexResponse
	TotalResults int64
}

type IcecreamIndexResponse struct {
	Id                    *int64    `json:"icecream_id"`
	Name                  *string   `json:"name,omitempty"`
	ImageClosed           *string   `json:"image_closed,omitempty"`
	ImageOpen             *string   `json:"image_open,omitempty"`
	Description           *string   `json:"description,omitempty"`
	Story                 *string   `json:"story,omitempty"`
	AllergyInfo           *string   `json:"allergy_info,omitempty"`
	SourcingValues        *[]string `json:"sourcing_values,omitempty"`
	Ingredients           *[]string `json:"ingredients,omitempty"`
	DietaryCertifications *string   `json:"dietary_certifications,omitempty"`
	ProductId             *string   `json:"product_id,omitempty"`
	//11
}
