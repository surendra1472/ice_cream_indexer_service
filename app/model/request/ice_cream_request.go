package request

import (
	"ic-indexer-service/app/utils"
	"time"
)

// Get icecream Request
// swagger:model icecreamRequest
type IcecreamFilter struct {
	// a icecream request
	// in: body
	Id        int64  `schema:"id" url:"id,omitempty"`
	Name      string `schema:"name"  url:"name,omitempty"`
	ProductId string `schema:"product_id"  url:"product_id,omitempty"`
}

// Delete icecream Request
// swagger:model icecreamDeleteRequest
type IcecreamDelete struct {
	// a icecream delete request
	// in: body
	ProductId string `schema:"product_id" validate:"required" url:"product_id,omitempty"`
}

type IcecreamIndexRequest struct {
	Id                    CusInt64       `json:"id" validate:"gt=0"`
	Name                  CusString      `json:"name,omitempty"`
	ImageClosed           CusString      `json:"image_closed,omitempty"`
	ImageOpen             CusString      `json:"image_open,omitempty"`
	Description           CusString      `json:"description,omitempty"`
	Story                 CusString      `json:"story,omitempty"`
	AllergyInfo           CusString      `json:"allergy_info,omitempty"`
	SourcingValues        CusArrayString `json:"sourcing_values,omitempty"`
	Ingredients           CusArrayString `json:"ingredients,omitempty"`
	DietaryCertifications CusString      `json:"dietary_certifications,omitempty"`
	ProductId             CusString      `json:"product_id,omitempty"`
}

// SaveOrUpdate icecream Request
// swagger:model icecreamClientRequest
type IcecreamClientRequest struct {
	// an icecream request
	// in: body
	Id                    int64    `json:"id" validate:"gt=0"`
	Name                  string   `json:"name,omitempty"`
	ImageClosed           string   `json:"image_closed,omitempty"`
	ImageOpen             string   `json:"image_open,omitempty"`
	Description           string   `json:"description,omitempty"`
	Story                 string   `json:"story,omitempty"`
	AllergyInfo           string   `json:"allergy_info,omitempty"`
	SourcingValues        []string `json:"sourcing_values,omitempty"`
	Ingredients           []string `json:"ingredients,omitempty"`
	DietaryCertifications string   `json:"dietary_certifications,omitempty"`
	ProductId             string   `json:"product_id,omitempty"`
	IsDeleted             bool     `json:"is_Deleted,omitempty"`
}

// Icecream Generic Response
//swagger:response genericResponse
type IcecreamGenericResponse struct {
	// in: body
	Body GenericModel
}

// Icecream Generic Model
// swagger:model genericModel
type GenericModel struct {
	Data      string `json:"data"`
	Code      int    `json:"code"`
	RequestId string `json:"request_id"`
	Error     string `json:"error"`
}

// Icecream Generic Error
// swagger:model genericError
type GenericError struct {
	Data      string `json:"data"`
	Code      int    `json:"code"`
	RequestId string `json:"request_id"`
	Error     string `json:"error"`
}

type CusTime struct {
	Value *time.Time
	Set   bool
}

type CusArrayString struct {
	Value *[]string
	Set   bool
}

type CusInt64 struct {
	Value *int64
	Set   bool
}

type CusString struct {
	Value *string
	Set   bool
}

func (x *CusString) UnmarshalJSON(data []byte) error {
	return utils.UnmarshalJSON(data, new(string), &x.Value, &x.Set)
}

func (x *CusInt64) UnmarshalJSON(data []byte) error {
	return utils.UnmarshalJSON(data, new(int64), &x.Value, &x.Set)
}
func (x *CusTime) UnmarshalJSON(data []byte) error {
	return utils.UnmarshalJSON(data, new(time.Time), &x.Value, &x.Set)
}

func (x *CusArrayString) UnmarshalJSON(data []byte) error {
	return utils.UnmarshalJSON(data, new([]string), &x.Value, &x.Set)
}
