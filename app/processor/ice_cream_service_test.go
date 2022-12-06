package processor

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	daMock "ic-indexer-service/app/api/dataaccessor/mocks"
	"ic-indexer-service/app/model/bo"
	"ic-indexer-service/app/model/request"
	"ic-indexer-service/app/model/response"
	"ic-indexer-service/app/processor/mocks"

	"testing"
)

func TestPartialUpdateCreateSuccess(t *testing.T) {
	dasMock := new(daMock.IceCreamDataAccessor)
	uhMock := new(mocks.UpdateHandler)
	is := GetNewIcecreamIndexerService(dasMock, uhMock)

	uhMock.On("UpdateIcecreamDetails", nil, mock.Anything, mock.Anything).Return(nil).Times(1)
	dasMock.On("GetIcecreamByKey", nil, mock.Anything).Return(bo.ESIcecream{}, nil).Times(1)
	dasMock.On("CreateOrReplaceIcecream", mock.Anything, mock.Anything).Return(nil).Times(1)
	productId := "test_1"
	err := is.PartialUpdate(nil, request.IcecreamIndexRequest{ProductId: request.CusString{Value: &productId, Set: true}})
	assert.Nil(t, err)

	dasMock.AssertNumberOfCalls(t, "GetIcecreamByKey", 1)
	uhMock.AssertNumberOfCalls(t, "UpdateIcecreamDetails", 1)
	dasMock.AssertNumberOfCalls(t, "CreateOrReplaceIcecream", 1)

}

func TestGetError(t *testing.T) {
	dasMock := new(daMock.IceCreamDataAccessor)
	uhMock := new(mocks.UpdateHandler)
	is := GetNewIcecreamIndexerService(dasMock, uhMock)

	dasMock.On("GetIcecreamByKey", nil, mock.Anything).Return(bo.ESIcecream{}, errors.New("error")).Times(1)

	err := is.PartialUpdate(nil, request.IcecreamIndexRequest{})
	assert.NotNil(t, err)
	dasMock.AssertNumberOfCalls(t, "GetIcecreamByKey", 1)
}

func TestGetSuccess(t *testing.T) {

	dasMock := new(daMock.IceCreamDataAccessor)
	uhMock := new(mocks.UpdateHandler)
	is := GetNewIcecreamIndexerService(dasMock, uhMock)

	testInt := int64(123)

	dasMock.On("GetIcecreamByKey", nil, mock.Anything).Return(bo.ESIcecream{Id: &testInt}, nil).Times(1)
	uhMock.On("UpdateIcecreamDetails", nil, mock.Anything, mock.Anything).Return(errors.New("error")).Times(1)
	product_id := "test_1"
	err := is.PartialUpdate(nil, request.IcecreamIndexRequest{ProductId: request.CusString{Value: &product_id, Set: true}})
	assert.NotNil(t, err)

	dasMock.AssertNumberOfCalls(t, "GetIcecreamByKey", 1)
	uhMock.AssertNumberOfCalls(t, "UpdateIcecreamDetails", 1)

}

func TestErrWhileCreate(t *testing.T) {

	dasMock := new(daMock.IceCreamDataAccessor)
	uhMock := new(mocks.UpdateHandler)
	is := GetNewIcecreamIndexerService(dasMock, uhMock)

	dasMock.On("GetIcecreamByKey", nil, mock.Anything).Return(bo.ESIcecream{}, nil).Times(1)
	uhMock.On("UpdateIcecreamDetails", nil, mock.Anything, mock.Anything).Return(nil).Times(1)
	dasMock.On("CreateOrReplaceIcecream", nil, mock.Anything).Return(errors.New("error")).Times(1)
	product_id := "test_1"
	err := is.PartialUpdate(nil, request.IcecreamIndexRequest{ProductId: request.CusString{Value: &product_id, Set: true}})
	assert.NotNil(t, err)

	dasMock.AssertNumberOfCalls(t, "GetIcecreamByKey", 1)
	uhMock.AssertNumberOfCalls(t, "UpdateIcecreamDetails", 1)
	dasMock.AssertNumberOfCalls(t, "CreateOrReplaceIcecream", 1)

}

func TestConstructor(t *testing.T) {
	dasMock := new(daMock.IceCreamDataAccessor)
	uhMock := new(mocks.UpdateHandler)
	is := GetNewIcecreamIndexerService(dasMock, uhMock)
	assert.NotNil(t, is)
}

func TestSearchSuccess(t *testing.T) {
	dasMock := new(daMock.IceCreamDataAccessor)
	uhMock := new(mocks.UpdateHandler)

	is := GetNewIcecreamIndexerService(dasMock, uhMock)

	params := request.IcecreamFilter{}

	dasMock.On("GetIcecreams", context.TODO(), mock.Anything, mock.Anything).Return(response.BulkIcecreamIndexResponse{}, nil).Times(1)

	response, err := is.Search(context.TODO(), params)
	dasMock.AssertNumberOfCalls(t, "GetIcecreams", 1)
	assert.Nil(t, err)
	assert.NotNil(t, response)
}

func TestDeleteSuccess(t *testing.T) {
	dasMock := new(daMock.IceCreamDataAccessor)
	uhMock := new(mocks.UpdateHandler)

	is := GetNewIcecreamIndexerService(dasMock, uhMock)

	params := request.IcecreamDelete{}

	dasMock.On("DeleteIcecream", context.TODO(), mock.Anything).Return(nil).Times(1)

	err := is.DeleteIcecream(context.TODO(), params)
	dasMock.AssertNumberOfCalls(t, "DeleteIcecream", 1)
	assert.Nil(t, err)
}
